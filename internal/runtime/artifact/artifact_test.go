package artifact

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func sha(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}

func sampleManifest(t *testing.T) Manifest {
	t.Helper()
	m := Manifest{
		ApplicationID: "org.example.ledger",
		Groups: []Group{
			{ID: "runtime-launch-metadata", Refs: []Ref{
				{ID: "launch.json", Kind: "metadata", SHA256: sha("launch-bytes"), Size: 12, Required: true},
			}},
			{ID: "local-execution-artifacts", Refs: []Ref{
				{ID: "payload.bin", Kind: "execution-artifacts", SHA256: sha("payload-bytes"), Size: 13, Required: false},
			}},
		},
	}
	data, _ := json.Marshal(m)
	parsed, err := ParseManifest(data)
	if err != nil {
		t.Fatalf("ParseManifest: %v", err)
	}
	return parsed
}

func TestParseManifestValidates(t *testing.T) {
	if _, err := ParseManifest([]byte(`{"application_id":"bad id","groups":[]}`)); err == nil {
		t.Fatalf("invalid application id must fail")
	}
	if _, err := ParseManifest([]byte(`{"application_id":"org.example.a","groups":[]}`)); err == nil {
		t.Fatalf("empty groups must fail")
	}
	bad := `{"application_id":"org.example.a","groups":[{"id":"g","refs":[{"id":"r","sha256":"nothex"}]}]}`
	if _, err := ParseManifest([]byte(bad)); err == nil {
		t.Fatalf("non-hex digest must fail")
	}
}

func TestCacheNamespaceAndKey(t *testing.T) {
	m := sampleManifest(t)
	if ns := m.CacheNamespace("runtime-launch-metadata"); ns != "org.example.ledger.runtime-launch-metadata" {
		t.Fatalf("unexpected namespace: %q", ns)
	}
}

func TestPlanIsDryRunAndReportsCacheKeys(t *testing.T) {
	m := sampleManifest(t)
	cacheDir := t.TempDir()
	cache, err := NewCache(cacheDir)
	if err != nil {
		t.Fatalf("NewCache: %v", err)
	}
	acq, _ := NewAcquirer(cache, nil)

	plan := acq.Plan(m)
	if !plan.DryRun || plan.NetworkEnabled || plan.HostRootMutated {
		t.Fatalf("plan must be a dry run with no network or host mutation: %#v", plan)
	}
	if len(plan.Artifacts) != 2 || plan.RequiredCount != 1 || plan.OptionalCount != 1 {
		t.Fatalf("unexpected plan artifacts: %#v", plan)
	}
	// allRefs sorts by group then ref id, so local-execution sorts before
	// runtime-launch-metadata. Locate the launch metadata artifact by ref id.
	var launch *PlannedArtifact
	for i := range plan.Artifacts {
		if plan.Artifacts[i].RefID == "launch.json" {
			launch = &plan.Artifacts[i]
		}
	}
	if launch == nil || launch.CacheKey != sha("launch-bytes") || launch.AlreadyCached || !launch.Required {
		t.Fatalf("unexpected launch metadata artifact: %#v", launch)
	}
	// A dry run must not have written anything into the cache root.
	entries, _ := os.ReadDir(cacheDir)
	if len(entries) != 0 {
		t.Fatalf("dry run must not stage files: %#v", entries)
	}
}

func TestAcquireStagesFromFixtureAndVerifies(t *testing.T) {
	m := sampleManifest(t)

	fixtureDir := t.TempDir()
	// Fixture files are addressed by digest.
	writeFixture(t, fixtureDir, sha("launch-bytes"), "launch-bytes")
	writeFixture(t, fixtureDir, sha("payload-bytes"), "payload-bytes")
	source, err := NewLocalFixtureSource(fixtureDir)
	if err != nil {
		t.Fatalf("NewLocalFixtureSource: %v", err)
	}

	cache, _ := NewCache(t.TempDir())
	acq, _ := NewAcquirer(cache, source)

	plan, err := acq.Acquire(m)
	if err != nil {
		t.Fatalf("Acquire: %v", err)
	}
	if plan.DryRun || plan.HostRootMutated {
		t.Fatalf("acquire plan flags wrong: %#v", plan)
	}
	if len(plan.StagedKeys) != 2 {
		t.Fatalf("expected 2 staged keys: %#v", plan.StagedKeys)
	}
	if !cache.Has(m.CacheNamespace("runtime-launch-metadata"), sha("launch-bytes")) {
		t.Fatalf("launch metadata not cached")
	}
	// Re-acquire is idempotent (already cached).
	plan2, err := acq.Acquire(m)
	if err != nil || len(plan2.StagedKeys) != 2 {
		t.Fatalf("re-acquire must be idempotent: %#v err=%v", plan2, err)
	}
}

func TestAcquireBlocksOnDigestMismatch(t *testing.T) {
	m := sampleManifest(t)

	fixtureDir := t.TempDir()
	// Store WRONG content under the expected digest name: the fetched bytes will
	// not hash to the ref digest, so staging must be blocked.
	writeFixture(t, fixtureDir, sha("launch-bytes"), "tampered-content")
	source, _ := NewLocalFixtureSource(fixtureDir)

	cache, _ := NewCache(t.TempDir())
	acq, _ := NewAcquirer(cache, source)

	if _, err := acq.Acquire(m); err == nil {
		t.Fatalf("digest mismatch must block staging")
	}
	if cache.Has(m.CacheNamespace("runtime-launch-metadata"), sha("launch-bytes")) {
		t.Fatalf("mismatched artifact must not be cached")
	}
}

func TestNetworkSourceDisabled(t *testing.T) {
	source := NewNetworkSource()
	if _, err := source.Fetch("ns", Ref{ID: "r", SHA256: sha("x")}); !errors.Is(err, ErrNetworkDisabled) {
		t.Fatalf("network source must be disabled, got %v", err)
	}
}

func TestCacheRefusesEscapeAndBadDigest(t *testing.T) {
	cache, _ := NewCache(t.TempDir())
	if err := cache.Put("../escape", sha("x"), []byte("x")); err == nil {
		t.Fatalf("namespace escape must be refused")
	}
	if err := cache.Put("ns", "not-a-digest", []byte("x")); err == nil {
		t.Fatalf("bad digest must be refused")
	}
	if _, err := NewCache(string(os.PathSeparator)); err == nil {
		t.Fatalf("filesystem root cache must be refused")
	}
}

func TestStageFromFixtureWritesReceiptUnderCacheRoot(t *testing.T) {
	manifest := sampleManifest(t)
	fixtureDir := t.TempDir()
	writeFixture(t, fixtureDir, sha("launch-bytes"), "launch-bytes")
	writeFixture(t, fixtureDir, sha("payload-bytes"), "payload-bytes")
	cacheRoot := t.TempDir()

	receipt, err := StageFromFixture(StageRequest{Manifest: manifest, CacheRoot: cacheRoot, FixtureDir: fixtureDir})
	if err != nil {
		t.Fatalf("StageFromFixture: %v", err)
	}
	if receipt.SchemaVersion != "xnix.runtime.artifact_stage_receipt.v1" ||
		receipt.RecordType != "compatibility-artifact-stage-receipt" ||
		receipt.Source != "go-runtime-local-fixture-artifact-staging" ||
		receipt.ApplicationID != manifest.ApplicationID ||
		receipt.RelativePath != "artifact-ledger/receipts/org.example.ledger.json" ||
		receipt.SHA256 == "" ||
		receipt.Plan.ApplicationID != manifest.ApplicationID ||
		len(receipt.Plan.StagedKeys) != 2 {
		t.Fatalf("unexpected stage receipt: %#v", receipt)
	}
	if !receipt.RuntimeOwned ||
		!receipt.GoRuntimeBacked ||
		receipt.KDEPolicyOwner ||
		receipt.CacheRootPathExposed ||
		receipt.FixtureRootPathExposed ||
		receipt.NetworkRequired ||
		receipt.NetworkFetchEnabled ||
		receipt.PackageManagerInvoked ||
		receipt.HostRootModified ||
		receipt.PrivilegedContainerRequired ||
		receipt.BackendLaunchEnabled ||
		receipt.BackendDetailsExposed {
		t.Fatalf("unexpected stage receipt safety flags: %#v", receipt)
	}
	if _, err := os.Stat(filepath.Join(cacheRoot, filepath.FromSlash(receipt.RelativePath))); err != nil {
		t.Fatalf("stage receipt was not written under cache root: %v", err)
	}
	if filepath.IsAbs(receipt.RelativePath) {
		t.Fatalf("stage receipt must expose only a relative path: %#v", receipt)
	}
}

func TestStageFromFixtureBlocksDigestMismatch(t *testing.T) {
	manifest := sampleManifest(t)
	fixtureDir := t.TempDir()
	writeFixture(t, fixtureDir, sha("launch-bytes"), "tampered-content")
	_, err := StageFromFixture(StageRequest{Manifest: manifest, CacheRoot: t.TempDir(), FixtureDir: fixtureDir})
	if err == nil {
		t.Fatalf("StageFromFixture must block digest mismatches")
	}
}

func TestStageFromFixtureRequiresExplicitRoots(t *testing.T) {
	manifest := sampleManifest(t)
	if _, err := StageFromFixture(StageRequest{Manifest: manifest, FixtureDir: t.TempDir()}); err == nil {
		t.Fatalf("missing cache root must be rejected")
	}
	if _, err := StageFromFixture(StageRequest{Manifest: manifest, CacheRoot: t.TempDir()}); err == nil {
		t.Fatalf("missing fixture root must be rejected")
	}
}

func writeFixture(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
}
