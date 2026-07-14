package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNewPlanPresentsCompatibilityAppAsNormalDesktopApp(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc", ".XLS"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	if plan.ApplicationID != "org.example.ledger" {
		t.Fatalf("ApplicationID = %q", plan.ApplicationID)
	}
	if plan.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("DesktopFile = %q", plan.DesktopFile)
	}
	if got, want := plan.LaunchCommand, []string{"xnix-compat-launch", "--app", "org.example.ledger", "%U"}; !sameStrings(got, want) {
		t.Fatalf("LaunchCommand = %#v, want %#v", got, want)
	}
	if got, want := plan.MIMETypes, []string{"application/x-xnix-abc", "application/x-xnix-xls"}; !sameStrings(got, want) {
		t.Fatalf("MIMETypes = %#v, want %#v", got, want)
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		t.Fatalf("ValidateSafeForDesktop returned error: %v", err)
	}
}

func TestPlanJSONHidesBackendTerminology(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.notepad",
		Name:                "Example Notepad",
		Icon:                "accessories-text-editor",
		Mode:                "wine",
		SupportedExtensions: []string{".txt"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	encoded, err := json.Marshal(plan)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("plan JSON exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestRenderDesktopEntryUsesManagedRuntimeLauncher(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	entry, err := plan.RenderDesktopEntry()
	if err != nil {
		t.Fatalf("RenderDesktopEntry returned error: %v", err)
	}
	required := []string{
		"[Desktop Entry]\n",
		"Type=Application\n",
		"Name=Example Ledger\n",
		"Exec=xnix-compat-launch --app org.example.ledger %U\n",
		"Icon=office-chart-area\n",
		"StartupWMClass=xnix-org.example.ledger\n",
		"X-Xnix-ApplicationId=org.example.ledger\n",
		"X-Xnix-RuntimeOwned=true\n",
		"MimeType=application/x-xnix-abc;application/x-xnix-xls;\n",
	}
	for _, fragment := range required {
		if !strings.Contains(entry, fragment) {
			t.Fatalf("desktop entry missing %q in:\n%s", fragment, entry)
		}
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system"} {
		if strings.Contains(strings.ToLower(entry), forbidden) {
			t.Fatalf("desktop entry exposes forbidden term %q: %s", forbidden, entry)
		}
	}
}

func TestRecipeValidationRejectsUnsafeIdentityInput(t *testing.T) {
	cases := []Recipe{
		{ID: "not-reverse-dns", Name: "Example", Icon: "icon", Mode: "automatic"},
		{ID: "org.example.app", Name: "Line\nBreak", Icon: "icon", Mode: "automatic"},
		{ID: "org.example.app", Name: "Example", Icon: "icon", Mode: "unknown"},
		{ID: "org.example.app", Name: "Example", Icon: "icon", Mode: "automatic", SupportedExtensions: []string{"bad"}},
	}
	for _, recipe := range cases {
		if _, err := NewPlan(recipe); err == nil {
			t.Fatalf("NewPlan accepted invalid recipe: %#v", recipe)
		}
	}
}

func sameStrings(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
