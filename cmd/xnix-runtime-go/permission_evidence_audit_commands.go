package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/portal"
)

func runPermissionEvidenceAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("permission-evidence-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	stateRoot := flags.String("state-root", "", "optional Runtime state root used for read-only Portal receipt evidence")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return errors.New("permission-evidence-audit-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return errors.New("permission-evidence-audit-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return errors.New("permission-evidence-audit-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return errors.New("permission-evidence-audit-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}

	receiptStates := map[string]string{}
	receiptRefs := map[string][]string{}
	var malformed []string
	if *stateRoot != "" {
		requests, malformedTokens, err := portal.ReadRequests(*stateRoot)
		if err != nil {
			return err
		}
		malformed = malformedTokens
		summary := portal.Summarize(requests, plan.ApplicationID)
		receiptStates, receiptRefs = permissionAuditReceiptEvidence(summary)
	}

	preview, err := plan.PermissionEvidenceAuditPreview(receiptStates, receiptRefs, malformed)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

// permissionAuditReceiptEvidence turns a KDE-safe Portal permission summary into
// per-operation receipt states plus safe, operation-scoped receipt references.
func permissionAuditReceiptEvidence(summary portal.PermissionSummary) (map[string]string, map[string][]string) {
	states := map[string]string{}
	refs := map[string][]string{}
	assign := func(operations []string, state string) {
		for _, operation := range operations {
			states[operation] = state
			refs[operation] = []string{"portal-operation:" + operation}
		}
	}
	assign(summary.Granted, "granted")
	assign(summary.Pending, "pending")
	assign(summary.Denied, "denied")
	assign(summary.Expired, "expired")
	return states, refs
}
