package portal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// ReadRequests reads persisted Portal request records under a state root in
// read-only mode, for audit and summary previews. Unlike NewLedger it never
// creates the ledger directory and never performs a real Portal call. Records
// that fail to parse are skipped and their file stems returned as malformed
// identifiers so callers can surface them without failing the whole read.
func ReadRequests(stateRoot string) ([]*Request, []string, error) {
	if strings.TrimSpace(stateRoot) == "" {
		return nil, nil, errors.New("portal request read requires a state root")
	}
	dir := filepath.Join(stateRoot, "portal-requests")
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil, nil
		}
		return nil, nil, fmt.Errorf("list portal request records: %w", err)
	}
	var requests []*Request
	var malformed []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		stem := strings.TrimSuffix(entry.Name(), ".json")
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			malformed = append(malformed, stem)
			continue
		}
		var request Request
		if err := json.Unmarshal(data, &request); err != nil {
			malformed = append(malformed, stem)
			continue
		}
		requests = append(requests, &request)
	}
	sort.Slice(requests, func(i, j int) bool {
		return requests[i].HandleToken < requests[j].HandleToken
	})
	sort.Strings(malformed)
	return requests, malformed, nil
}
