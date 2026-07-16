package preview

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSafeDefaultsSatisfyInvariant(t *testing.T) {
	f := Safe()
	if !f.RuntimeOwned || !f.GoRuntimeBacked {
		t.Fatalf("Safe defaults must be runtime-owned and go-backed: %#v", f)
	}
	if !f.Safe() {
		t.Fatalf("Safe defaults must satisfy the invariant")
	}
	if err := f.Validate(); err != nil {
		t.Fatalf("Safe defaults must validate: %v", err)
	}
}

func TestValidateRejectsEachUnsafeFlag(t *testing.T) {
	base := Safe()

	mutate := map[string]func(*SafetyFlags){
		"kde_policy_owner":              func(f *SafetyFlags) { f.KDEPolicyOwner = true },
		"host_root_modified":            func(f *SafetyFlags) { f.HostRootModified = true },
		"network_required":              func(f *SafetyFlags) { f.NetworkRequired = true },
		"privileged_container_required": func(f *SafetyFlags) { f.PrivilegedContainerRequired = true },
		"backend_launch_enabled":        func(f *SafetyFlags) { f.BackendLaunchEnabled = true },
		"backend_details_exposed":       func(f *SafetyFlags) { f.BackendDetailsExposed = true },
	}
	for flag, apply := range mutate {
		f := base
		apply(&f)
		err := f.Validate()
		if err == nil {
			t.Fatalf("setting %s must violate the invariant", flag)
		}
		if !strings.Contains(err.Error(), flag) {
			t.Fatalf("error should name %s: %v", flag, err)
		}
		if f.Safe() {
			t.Fatalf("flags with %s set must not be Safe", flag)
		}
	}
}

func TestValidateRequiresRuntimeOwned(t *testing.T) {
	f := Safe()
	f.RuntimeOwned = false
	if err := f.Validate(); err == nil || !strings.Contains(err.Error(), "runtime_owned") {
		t.Fatalf("runtime_owned=false must violate the invariant: %v", err)
	}
}

func TestSafetyFlagsFlattenInJSON(t *testing.T) {
	// Embedding SafetyFlags must produce top-level keys, so adopting it does not
	// change a payload's JSON shape (only field ordering).
	type payload struct {
		Name string `json:"name"`
		SafetyFlags
	}
	data, err := json.Marshal(payload{Name: "x", SafetyFlags: Safe()})
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	for _, key := range []string{"runtime_owned", "go_runtime_backed", "kde_policy_owner", "host_root_modified", "backend_details_exposed"} {
		if _, ok := decoded[key]; !ok {
			t.Fatalf("embedded flag %q did not flatten to top level: %s", key, data)
		}
	}
	if decoded["runtime_owned"] != true || decoded["backend_details_exposed"] != false {
		t.Fatalf("unexpected flag values: %s", data)
	}
}

func TestHeaderValidate(t *testing.T) {
	ok := Header{SchemaVersion: "xnix.runtime.launch_readiness.v1", RequestType: "execution-readiness-preview", RuntimeMethod: "GetExecutionReadiness", Desktop: "KDE Plasma"}
	if err := ok.Validate(); err != nil {
		t.Fatalf("valid header rejected: %v", err)
	}

	bad := []Header{
		{SchemaVersion: "launch_readiness.v1", RequestType: "r", RuntimeMethod: "m"},                 // bad schema
		{SchemaVersion: "xnix.runtime.x.v1", RuntimeMethod: "m"},                                     // missing request_type
		{SchemaVersion: "xnix.runtime.x.v1", RequestType: "r"},                                       // missing runtime_method
		{SchemaVersion: "xnix.runtime.x.v1", RequestType: "r", RuntimeMethod: "m", Desktop: "GNOME"}, // wrong desktop
	}
	for i, h := range bad {
		if err := h.Validate(); err == nil {
			t.Fatalf("bad header %d should be rejected: %#v", i, h)
		}
	}
}
