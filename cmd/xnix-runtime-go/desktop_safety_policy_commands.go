package main

import (
	"encoding/json"
	"errors"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runDesktopSafetyPolicyPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return errors.New("desktop-safety-policy-preview does not accept arguments")
	}

	preview := appidentity.NewDesktopSafetyPolicyPreview()
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
