package appidentity

import (
	"testing"

	"xnix.local/xnix/internal/testversion"
)

func currentProjectVersion(t *testing.T) string {
	t.Helper()
	version, err := testversion.Current()
	if err != nil {
		t.Fatalf("read project version: %v", err)
	}
	return version
}
