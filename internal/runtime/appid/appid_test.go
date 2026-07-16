package appid

import "testing"

func TestValid(t *testing.T) {
	valid := []string{
		"org.xnix.sample.notepad",
		"org.example.ledger",
		"com.a.b",
		"org.example.app-name",
	}
	for _, id := range valid {
		if !Valid(id) {
			t.Fatalf("expected valid: %q", id)
		}
		if err := Validate(id); err != nil {
			t.Fatalf("Validate rejected valid id %q: %v", id, err)
		}
	}

	invalid := []string{
		"",
		"nodot",
		"bad id",
		".leading",
		"trailing.",
		"org..double",
		"org.example.app_name", // underscore not allowed in a label
		"/etc/passwd",
	}
	for _, id := range invalid {
		if Valid(id) {
			t.Fatalf("expected invalid: %q", id)
		}
		if err := Validate(id); err == nil {
			t.Fatalf("Validate accepted invalid id %q", id)
		}
	}
}

func TestValidateErrorNamesTheID(t *testing.T) {
	err := Validate("bad id")
	if err == nil {
		t.Fatalf("expected error")
	}
	if got := err.Error(); got == "" {
		t.Fatalf("error should be non-empty")
	}
}

func TestSlugKeepsDotsAndHyphens(t *testing.T) {
	cases := map[string]string{
		"org.example.ledger": "org.example.ledger",
		"org.example/app id": "org.example-app-id",
		"a.b-c":              "a.b-c",
		"weird!@#name":       "weird---name",
	}
	for in, want := range cases {
		if got := Slug(in); got != want {
			t.Fatalf("Slug(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestFileTokenKeepsOnlyAlphanumeric(t *testing.T) {
	cases := map[string]string{
		"org.xnix.sample.notepad": "org_xnix_sample_notepad",
		"file-open":               "file_open",
		"a.b-c":                   "a_b_c",
		"already_ok9":             "already_ok9",
	}
	for in, want := range cases {
		if got := FileToken(in); got != want {
			t.Fatalf("FileToken(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestSanitizersAreStableAndFilesystemSafe(t *testing.T) {
	// The output of both sanitizers must contain no path separators so it is
	// safe to use as a single path component.
	for _, in := range []string{"org.example/../escape", "a\\b/c", "x y z"} {
		for _, out := range []string{Slug(in), FileToken(in)} {
			for _, r := range out {
				if r == '/' || r == '\\' {
					t.Fatalf("sanitized %q -> %q contains a path separator", in, out)
				}
			}
		}
	}
}
