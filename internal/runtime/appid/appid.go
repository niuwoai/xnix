// Package appid is the canonical application-identifier boundary for the
// Compatibility Runtime. Application ids are reverse-DNS identifiers
// (e.g. org.xnix.sample.notepad); the same validation regex was copied
// verbatim into at least five packages, and several packages also carried
// near-duplicate helpers that sanitize an id into a filesystem-safe token with
// divergent rules. This package is the single reviewed, tested source for
// validation and for the two sanitization variants those call sites need.
package appid

import (
	"errors"
	"fmt"
	"regexp"
)

// pattern matches a reverse-DNS application identifier: alphanumeric labels
// separated by dots, with at least one dot.
var pattern = regexp.MustCompile(`^[A-Za-z0-9]+(\.[A-Za-z0-9-]+)+$`)

// Valid reports whether id is a well-formed reverse-DNS application identifier.
func Valid(id string) bool {
	return pattern.MatchString(id)
}

// Validate returns an error when id is not a reverse-DNS application identifier.
func Validate(id string) error {
	if !Valid(id) {
		return fmt.Errorf("application id %q must be a reverse-DNS identifier", id)
	}
	return nil
}

// ErrInvalid is a sentinel for callers that prefer errors.Is over Validate's
// formatted message.
var ErrInvalid = errors.New("application id must be a reverse-DNS identifier")

// Slug returns a filesystem-safe form of value that keeps ASCII letters,
// digits, dots, and hyphens and replaces every other rune with a hyphen. Use it
// for cache namespaces and state namespaces that may keep dots and hyphens.
func Slug(value string) string {
	return sanitize(value, '-', true)
}

// FileToken returns a filesystem-safe form of value that keeps only ASCII
// letters and digits and replaces every other rune with an underscore. Use it
// for handle tokens and record keys that must not contain separators.
func FileToken(value string) string {
	return sanitize(value, '_', false)
}

// sanitize replaces disallowed runes with repl. When keepDotDash is true, '.'
// and '-' are preserved.
func sanitize(value string, repl rune, keepDotDash bool) string {
	out := make([]rune, 0, len(value))
	for _, r := range value {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
			out = append(out, r)
		case keepDotDash && (r == '.' || r == '-'):
			out = append(out, r)
		default:
			out = append(out, repl)
		}
	}
	return string(out)
}
