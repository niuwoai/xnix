package environment

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Lifecycle persists compatibility-environment records under a controlled state
// root and enforces valid transitions between lifecycle states.
type Lifecycle struct {
	root string
	dir  string
}

// New opens (creating if needed) a lifecycle store under the given state root.
// The root must be a caller-controlled directory (e.g. a test state root); the
// store never writes outside it and never touches the host root.
func New(stateRoot string) (*Lifecycle, error) {
	if stateRoot == "" {
		return nil, errors.New("environment lifecycle requires a state root")
	}
	abs, err := filepath.Abs(stateRoot)
	if err != nil {
		return nil, fmt.Errorf("resolve state root: %w", err)
	}
	info, err := os.Stat(abs)
	if err != nil {
		return nil, fmt.Errorf("state root must exist: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("state root %q is not a directory", abs)
	}
	dir := filepath.Join(abs, "environments")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, fmt.Errorf("initialize environment store: %w", err)
	}
	return &Lifecycle{root: abs, dir: dir}, nil
}

func recordKey(applicationID string, profile Profile) string {
	return sanitize(applicationID) + "__" + sanitize(string(profile))
}

func sanitize(value string) string {
	sanitized := make([]rune, 0, len(value))
	for _, r := range value {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '-':
			sanitized = append(sanitized, r)
		default:
			sanitized = append(sanitized, '_')
		}
	}
	return string(sanitized)
}

func (l *Lifecycle) recordPath(applicationID string, profile Profile) string {
	return filepath.Join(l.dir, recordKey(applicationID, profile)+".json")
}

// Get returns the persisted record, or a fresh missing-state record if none
// exists yet.
func (l *Lifecycle) Get(applicationID string, profile Profile) (Record, error) {
	path := l.recordPath(applicationID, profile)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return newRecord(applicationID, profile)
		}
		return Record{}, fmt.Errorf("read environment record: %w", err)
	}
	var record Record
	if err := json.Unmarshal(data, &record); err != nil {
		return Record{}, fmt.Errorf("environment record is corrupt: %w", err)
	}
	if !record.State.valid() {
		return Record{}, fmt.Errorf("environment record has invalid state %q", string(record.State))
	}
	return record, nil
}

func (l *Lifecycle) save(record Record) error {
	// Safety invariants that must always hold for a preview-mode lifecycle.
	record.LaunchEnabled = false
	record.BackendStarted = false
	record.HostRootModified = false
	record.BackendDetailsShown = false
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(l.recordPath(record.ApplicationID, record.Profile), data, 0o600)
}

func (l *Lifecycle) transition(applicationID string, profile Profile, to State, mutate func(*Record)) (Record, error) {
	record, err := l.Get(applicationID, profile)
	if err != nil {
		return Record{}, err
	}
	if record.State == to && to != StateRepairRequired {
		return record, nil // idempotent for non-repair states
	}
	if !canTransition(record.State, to) {
		return Record{}, fmt.Errorf("invalid transition %s -> %s for %s/%s", record.State, to, applicationID, profile)
	}
	record.State = to
	if mutate != nil {
		mutate(&record)
	}
	if err := l.save(record); err != nil {
		return Record{}, err
	}
	return record, nil
}

// Plan moves a missing environment to planned.
func (l *Lifecycle) Plan(applicationID string, profile Profile) (Record, error) {
	return l.transition(applicationID, profile, StatePlanned, func(r *Record) {
		r.RepairHints = []string{}
		r.BlockReason = ""
	})
}

// Stage moves a planned or repaired environment to staged.
func (l *Lifecycle) Stage(applicationID string, profile Profile) (Record, error) {
	return l.transition(applicationID, profile, StateStaged, func(r *Record) {
		r.RepairHints = []string{}
	})
}

// SatisfyGate records that a required preflight gate is satisfied. It does not
// change lifecycle state; call MarkReady once all gates are satisfied.
func (l *Lifecycle) SatisfyGate(applicationID string, profile Profile, gate string) (Record, error) {
	if !isRequiredGate(gate) {
		return Record{}, fmt.Errorf("unknown environment gate %q", gate)
	}
	record, err := l.Get(applicationID, profile)
	if err != nil {
		return Record{}, err
	}
	if record.State != StateStaged && record.State != StateReady {
		return Record{}, fmt.Errorf("gates can only be satisfied while staged or ready, not %s", record.State)
	}
	for _, g := range record.SatisfiedGates {
		if g == gate {
			return record, nil
		}
	}
	record.SatisfiedGates = append(record.SatisfiedGates, gate)
	sort.Strings(record.SatisfiedGates)
	if err := l.save(record); err != nil {
		return Record{}, err
	}
	return record, nil
}

// MarkReady moves a staged environment to ready, but only when every required
// gate is satisfied. Readiness never enables launch.
func (l *Lifecycle) MarkReady(applicationID string, profile Profile) (Record, error) {
	record, err := l.Get(applicationID, profile)
	if err != nil {
		return Record{}, err
	}
	if pending := record.PendingGates(); len(pending) > 0 {
		return Record{}, fmt.Errorf("cannot mark ready; pending gates: %s", strings.Join(pending, ", "))
	}
	return l.transition(applicationID, profile, StateReady, nil)
}

// FlagRepair moves an environment to repair-required with hints. Hints describe
// what a user or Runtime should do; no repair is executed.
func (l *Lifecycle) FlagRepair(applicationID string, profile Profile, hints []string) (Record, error) {
	if len(hints) == 0 {
		return Record{}, errors.New("repair requires at least one hint")
	}
	return l.transition(applicationID, profile, StateRepairRequired, func(r *Record) {
		r.RepairHints = append([]string(nil), hints...)
		// A repair invalidates any satisfied gates.
		r.SatisfiedGates = []string{}
	})
}

// Block holds an environment closed with a reason.
func (l *Lifecycle) Block(applicationID string, profile Profile, reason string) (Record, error) {
	if strings.TrimSpace(reason) == "" {
		return Record{}, errors.New("block requires a reason")
	}
	return l.transition(applicationID, profile, StateBlocked, func(r *Record) {
		r.BlockReason = reason
	})
}

// Retire decommissions an environment. A retired environment holds no satisfied
// gates, repair hints, or block reason, and can be reactivated by planning it
// again. It launches nothing and mutates no host state.
func (l *Lifecycle) Retire(applicationID string, profile Profile) (Record, error) {
	return l.transition(applicationID, profile, StateRetired, func(r *Record) {
		r.SatisfiedGates = []string{}
		r.RepairHints = []string{}
		r.BlockReason = ""
	})
}

// List returns every persisted environment record, sorted by key.
func (l *Lifecycle) List() ([]Record, error) {
	entries, err := os.ReadDir(l.dir)
	if err != nil {
		return nil, fmt.Errorf("list environments: %w", err)
	}
	var out []Record
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		data, readErr := os.ReadFile(filepath.Join(l.dir, entry.Name()))
		if readErr != nil {
			return nil, readErr
		}
		var record Record
		if err := json.Unmarshal(data, &record); err != nil {
			return nil, fmt.Errorf("environment record %s is corrupt: %w", entry.Name(), err)
		}
		out = append(out, record)
	}
	sort.Slice(out, func(i, j int) bool {
		return recordKey(out[i].ApplicationID, out[i].Profile) < recordKey(out[j].ApplicationID, out[j].Profile)
	})
	return out, nil
}
