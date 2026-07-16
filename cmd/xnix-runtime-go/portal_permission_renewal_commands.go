package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/portal"
)

func runPortalPermissionRenewalPreview(args []string, stdout io.Writer) error {
	plan, receipts, malformed, err := parsePortalPermissionRenewalPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := plan.PortalPermissionRenewalPreview(receipts, malformed)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parsePortalPermissionRenewalPreviewSource(args []string) (appidentity.Plan, []appidentity.PortalPermissionRenewalReceipt, []string, error) {
	flags := flag.NewFlagSet("portal-permission-renewal-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	stateRoot := flags.String("state-root", "", "optional Runtime state root used for read-only Portal receipt evidence")
	expiring := flags.String("expiring", "", "comma-separated Portal operations that should be treated as expiring soon")
	if err := flags.Parse(args); err != nil {
		return appidentity.Plan{}, nil, nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Plan{}, nil, nil, errors.New("portal-permission-renewal-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Plan{}, nil, nil, errors.New("portal-permission-renewal-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Plan{}, nil, nil, errors.New("portal-permission-renewal-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Plan{}, nil, nil, errors.New("portal-permission-renewal-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return appidentity.Plan{}, nil, nil, err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return appidentity.Plan{}, nil, nil, err
	}

	var requests []*portal.Request
	var malformed []string
	if *stateRoot != "" {
		requests, malformed, err = portal.ReadRequests(*stateRoot)
		if err != nil {
			return appidentity.Plan{}, nil, nil, err
		}
	}
	receipts := portalPermissionRenewalReceipts(requests, plan.ApplicationID, csvSet(*expiring))
	return plan, receipts, malformed, nil
}

func portalPermissionRenewalReceipts(requests []*portal.Request, applicationID string, expiring map[string]bool) []appidentity.PortalPermissionRenewalReceipt {
	latest := map[string]appidentity.PortalPermissionRenewalReceipt{}
	for _, request := range requests {
		if request == nil || request.ApplicationID != applicationID {
			continue
		}
		state := portalPermissionRenewalStateFromRequest(*request)
		operation := strings.TrimSpace(request.Operation)
		if operation == "" {
			continue
		}
		latest[operation] = appidentity.PortalPermissionRenewalReceipt{
			Operation:    operation,
			State:        state,
			ExpiringSoon: expiring[operation],
			EvidenceID:   "portal-operation:" + operation,
		}
	}
	receipts := make([]appidentity.PortalPermissionRenewalReceipt, 0, len(latest))
	for _, receipt := range latest {
		receipts = append(receipts, receipt)
	}
	return receipts
}

func portalPermissionRenewalStateFromRequest(request portal.Request) string {
	switch request.State {
	case portal.StateGranted, portal.StateCompleted:
		return "granted"
	case portal.StateDenied:
		return "denied"
	case portal.StateCancelled:
		return "revoked"
	case portal.StateExpired:
		return "expired"
	default:
		return "pending"
	}
}

func csvSet(value string) map[string]bool {
	out := map[string]bool{}
	for _, part := range strings.Split(value, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		out[part] = true
	}
	return out
}
