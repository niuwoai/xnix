package artifact

import "testing"

func TestRequiredStagedAfterAcquire(t *testing.T) {
	m := sampleManifest(t)

	fixtureDir := t.TempDir()
	writeFixture(t, fixtureDir, sha("launch-bytes"), "launch-bytes")
	writeFixture(t, fixtureDir, sha("payload-bytes"), "payload-bytes")
	source, _ := NewLocalFixtureSource(fixtureDir)
	cache, _ := NewCache(t.TempDir())
	acq, _ := NewAcquirer(cache, source)

	plan, err := acq.Acquire(m)
	if err != nil {
		t.Fatalf("Acquire: %v", err)
	}
	if !plan.RequiredStaged() {
		t.Fatalf("all required artifacts staged after acquire: %#v", plan)
	}
}

func TestRequiredStagedFalseOnDryRun(t *testing.T) {
	m := sampleManifest(t)
	cache, _ := NewCache(t.TempDir())
	acq, _ := NewAcquirer(cache, nil)

	plan := acq.Plan(m)
	if plan.RequiredStaged() {
		t.Fatalf("a dry-run plan with nothing cached must not be staged: %#v", plan)
	}
}

func TestRequiredStagedIgnoresOptional(t *testing.T) {
	m := sampleManifest(t)
	// The sample manifest's only required artifact is the launch metadata;
	// payload is optional. Stage only the required artifact directly.
	cache, _ := NewCache(t.TempDir())
	if err := cache.Put(m.CacheNamespace("runtime-launch-metadata"), sha("launch-bytes"), []byte("launch-bytes")); err != nil {
		t.Fatalf("Put: %v", err)
	}
	acq, _ := NewAcquirer(cache, nil)

	plan := acq.Plan(m)
	// The required launch metadata is cached; the optional payload is not.
	if !plan.RequiredStaged() {
		t.Fatalf("required-only staging must satisfy RequiredStaged: %#v", plan)
	}
}
