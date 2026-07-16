package testversion

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Current returns the canonical repository version from the working directory.
func Current() (string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}
	return ReadFrom(workingDirectory)
}

// ReadFrom walks upward from start until it finds the repository VERSION file.
func ReadFrom(start string) (string, error) {
	directory, err := filepath.Abs(start)
	if err != nil {
		return "", fmt.Errorf("resolve version search root %q: %w", start, err)
	}

	for {
		contents, readErr := os.ReadFile(filepath.Join(directory, "VERSION"))
		if readErr == nil {
			version := strings.TrimSpace(string(contents))
			if version == "" {
				return "", fmt.Errorf("VERSION is empty under %s", directory)
			}
			return version, nil
		}
		if !os.IsNotExist(readErr) {
			return "", fmt.Errorf("read VERSION under %s: %w", directory, readErr)
		}

		parent := filepath.Dir(directory)
		if parent == directory {
			return "", fmt.Errorf("could not find VERSION from %s", start)
		}
		directory = parent
	}
}
