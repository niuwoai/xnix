package safety

import (
	"strings"
	"testing"
)

func TestSafeAcceptsUserFacingText(t *testing.T) {
	ok := []string{
		"Compatibility application is ready.",
		"Runtime backend lifecycle is modeled but blocked.",
		"org.xnix.sample.notepad",
		"xnix-org.example.ledger.desktop",
		"/usr/lib/systemd/system/xnix-compatd.service", // system path, not a host user path
		"KDE Plasma",
	}
	for _, text := range ok {
		if !Safe(text) {
			t.Fatalf("expected safe: %q -> %#v", text, Scan(text))
		}
		if err := Validate("field", text); err != nil {
			t.Fatalf("Validate rejected safe text %q: %v", text, err)
		}
	}
}

func TestScanCatchesEachCategory(t *testing.T) {
	cases := map[string]Category{
		"launch via wine backend":            CategoryBackendTerm,
		"start proton runtime":               CategoryBackendTerm,
		"qemu-system-x86_64 started":         CategoryBackendTerm,
		"C:\\Program Files\\app":             CategoryWindowsPath, // also windows-path; program files -> backend-term
		"opened /home/rocky/Documents/x.txt": CategoryHostPath,
		"read /Users/rocky/secret":           CategoryHostPath,
		"ran notepad.exe":                    CategoryExecutable,
		"api key sk-abcdef123456 leaked":     CategorySecret,
		"Authorization: Bearer abcdef123456": CategorySecret,
		"password is hunter2":                CategorySecret,
		"-----BEGIN RSA PRIVATE KEY-----":    CategoryPrivateKey,
	}
	for text, want := range cases {
		findings := Scan(text)
		found := false
		for _, f := range findings {
			if f.Category == want {
				found = true
			}
		}
		if !found {
			t.Fatalf("expected category %q for %q, got %#v", want, text, findings)
		}
		if Safe(text) {
			t.Fatalf("expected unsafe: %q", text)
		}
	}
}

func TestValidateNamesCategoriesWithoutLeakingValue(t *testing.T) {
	err := Validate("advice", "your token is sk-supersecretvalue123")
	if err == nil {
		t.Fatalf("expected error for secret content")
	}
	msg := err.Error()
	if !strings.Contains(msg, "secret") {
		t.Fatalf("error should name the category: %q", msg)
	}
	if strings.Contains(msg, "supersecretvalue") {
		t.Fatalf("error must not echo the offending value: %q", msg)
	}
}

func TestValidateLineRequiresSingleLine(t *testing.T) {
	if err := ValidateLine("id", "safe text"); err != nil {
		t.Fatalf("single-line safe text rejected: %v", err)
	}
	if err := ValidateLine("id", "line one\nline two"); err == nil {
		t.Fatalf("multi-line text must be rejected")
	}
}

func TestValidatePayloadCatchesNestedForbiddenContent(t *testing.T) {
	type inner struct {
		Note string `json:"note"`
	}
	type payload struct {
		Name    string   `json:"name"`
		Details inner    `json:"details"`
		Tags    []string `json:"tags"`
	}
	safe := payload{Name: "Ledger", Details: inner{Note: "ready"}, Tags: []string{"office"}}
	if err := ValidatePayload("payload", safe); err != nil {
		t.Fatalf("safe payload rejected: %v", err)
	}
	unsafe := payload{Name: "Ledger", Details: inner{Note: "using wine prefix at /home/rocky/.wine"}, Tags: []string{"office"}}
	if err := ValidatePayload("payload", unsafe); err == nil {
		t.Fatalf("nested forbidden content must be caught")
	}
}

func TestRedactReplacesForbiddenSpans(t *testing.T) {
	got := Redact("open /home/rocky/x with wine backend")
	if strings.Contains(got, "/home/rocky") || strings.Contains(strings.ToLower(got), "wine") {
		t.Fatalf("redaction left forbidden content: %q", got)
	}
	if !strings.Contains(got, Placeholder) {
		t.Fatalf("redaction should insert placeholder: %q", got)
	}
	// Redacted output must itself be safe.
	if !Safe(got) {
		t.Fatalf("redacted output is not safe: %q -> %#v", got, Scan(got))
	}
}

func TestScanIsDeterministicallyOrdered(t *testing.T) {
	findings := Scan("wine at /home/rocky with token secret")
	if len(findings) < 2 {
		t.Fatalf("expected multiple categories: %#v", findings)
	}
	for i := 1; i < len(findings); i++ {
		if findings[i-1].Category > findings[i].Category {
			t.Fatalf("findings not sorted: %#v", findings)
		}
	}
}
