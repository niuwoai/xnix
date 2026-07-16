// Package safety is the canonical KDE-facing / user-facing output safety
// boundary for the Compatibility Runtime. Every Runtime domain must keep raw
// backend commands, executable paths, compatibility storage paths, profile
// terminology, user file contents, secrets, tokens, private keys, and host
// paths out of user-visible output. That rule was previously reimplemented,
// with divergent term lists, in several packages; this package consolidates it
// into one reviewed, tested validator and redactor.
//
// The validator is content-only: it scans text (or a JSON-marshaled payload)
// for forbidden patterns and reports the matching categories without echoing
// the offending value, so an error never leaks the secret it caught.
package safety

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Category names a class of forbidden content.
type Category string

const (
	// CategoryBackendTerm covers compatibility backend implementation terms.
	CategoryBackendTerm Category = "backend-term"
	// CategoryHostPath covers user-home and OS-user host paths.
	CategoryHostPath Category = "host-path"
	// CategoryWindowsPath covers Windows drive-letter paths.
	CategoryWindowsPath Category = "windows-path"
	// CategorySecret covers secrets, tokens, and passwords.
	CategorySecret Category = "secret"
	// CategoryPrivateKey covers PEM private key material.
	CategoryPrivateKey Category = "private-key"
	// CategoryExecutable covers raw executable file references.
	CategoryExecutable Category = "executable"
)

type rule struct {
	category Category
	pattern  *regexp.Regexp
}

// rules is the ordered, canonical set of forbidden-content detectors. It is a
// superset of the previously scattered lists: backend terms plus host paths,
// Windows paths, secrets, private keys, and raw executables.
var rules = []rule{
	// Backend implementation terminology.
	{CategoryBackendTerm, regexp.MustCompile(`(?i)\bwine\b`)},
	{CategoryBackendTerm, regexp.MustCompile(`(?i)\.wine\b`)},
	{CategoryBackendTerm, regexp.MustCompile(`(?i)\bproton\b`)},
	{CategoryBackendTerm, regexp.MustCompile(`(?i)qemu-system`)},
	{CategoryBackendTerm, regexp.MustCompile(`(?i)\bwineprefix\b`)},
	{CategoryBackendTerm, regexp.MustCompile(`(?i)program files`)},
	{CategoryExecutable, regexp.MustCompile(`(?i)\.exe\b`)},
	// Host paths that identify a real user or machine.
	{CategoryHostPath, regexp.MustCompile(`(?i)/home/[^\s"]+`)},
	{CategoryHostPath, regexp.MustCompile(`(?i)/users/[^\s"]+`)},
	{CategoryHostPath, regexp.MustCompile(`/root/[^\s"]+`)},
	{CategoryWindowsPath, regexp.MustCompile(`(?i)[a-z]:\\`)},
	// Secrets, tokens, and credentials.
	{CategorySecret, regexp.MustCompile(`(?i)\bsecret\b`)},
	{CategorySecret, regexp.MustCompile(`(?i)\btoken\b`)},
	{CategorySecret, regexp.MustCompile(`(?i)\bpassword\b`)},
	{CategorySecret, regexp.MustCompile(`\bsk-[A-Za-z0-9]{8,}`)},
	{CategorySecret, regexp.MustCompile(`(?i)\bbearer\s+[A-Za-z0-9._-]{8,}`)},
	{CategoryPrivateKey, regexp.MustCompile(`(?i)-----begin [a-z0-9 ]*private key-----`)},
}

// Finding is one forbidden-content match. It records the category only, never
// the matched text, so surfacing findings cannot leak the offending value.
type Finding struct {
	Category Category `json:"category"`
}

// Scan returns the distinct categories of forbidden content found in text.
func Scan(text string) []Finding {
	seen := map[Category]bool{}
	for _, r := range rules {
		if !seen[r.category] && r.pattern.MatchString(text) {
			seen[r.category] = true
		}
	}
	cats := make([]Category, 0, len(seen))
	for c := range seen {
		cats = append(cats, c)
	}
	sort.Slice(cats, func(i, j int) bool { return cats[i] < cats[j] })
	findings := make([]Finding, 0, len(cats))
	for _, c := range cats {
		findings = append(findings, Finding{Category: c})
	}
	return findings
}

// Safe reports whether text contains no forbidden content.
func Safe(text string) bool {
	for _, r := range rules {
		if r.pattern.MatchString(text) {
			return false
		}
	}
	return true
}

// Validate returns an error naming the offending categories (never the value)
// when text contains forbidden content. label identifies the field for the
// error message.
func Validate(label, text string) error {
	findings := Scan(text)
	if len(findings) == 0 {
		return nil
	}
	cats := make([]string, 0, len(findings))
	for _, f := range findings {
		cats = append(cats, string(f.Category))
	}
	return fmt.Errorf("%s exposes forbidden content: %s", label, strings.Join(cats, ", "))
}

// ValidateLine is Validate plus a single-line requirement, for identity fields
// that must never contain embedded newlines.
func ValidateLine(label, text string) error {
	if strings.ContainsAny(text, "\r\n") {
		return fmt.Errorf("%s must be a single line", label)
	}
	return Validate(label, text)
}

// ValidatePayload marshals v to JSON and validates the encoded form. It is the
// canonical replacement for the per-package "validateNoBackendTerms" helpers:
// it catches forbidden content anywhere in a KDE-facing payload.
func ValidatePayload(label string, v any) error {
	encoded, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("encode %s for safety validation: %w", label, err)
	}
	return Validate(label, string(encoded))
}

// Placeholder is substituted for redacted spans.
const Placeholder = "[redacted]"

// Redact replaces every forbidden span in text with Placeholder. Use it for
// best-effort sanitization of text that may contain unsafe fragments; prefer
// Validate to reject unsafe content outright at the source.
func Redact(text string) string {
	out := text
	for _, r := range rules {
		out = r.pattern.ReplaceAllString(out, Placeholder)
	}
	return out
}
