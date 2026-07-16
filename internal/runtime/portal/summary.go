package portal

import "sort"

// PermissionSummary is a KDE-safe roll-up of an application's portal requests,
// for the Compatibility Center permission view and execution preflight. It
// reports the latest state per operation and exposes no host paths, backend
// details, or file contents.
type PermissionSummary struct {
	ApplicationID         string   `json:"application_id"`
	Granted               []string `json:"granted"`
	Pending               []string `json:"pending"`
	Denied                []string `json:"denied"`
	Expired               []string `json:"expired"`
	Total                 int      `json:"total"`
	AllResolved           bool     `json:"all_resolved"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

// Summarize rolls up the requests belonging to applicationID into a permission
// summary. When an operation has more than one request (e.g. a retry), the most
// recent request wins. Requests preserve creation order in a broker's List, so
// callers should pass that order.
func Summarize(requests []*Request, applicationID string) PermissionSummary {
	latest := map[string]*Request{}
	order := make([]string, 0)
	for _, r := range requests {
		if r == nil || r.ApplicationID != applicationID {
			continue
		}
		if _, seen := latest[r.Operation]; !seen {
			order = append(order, r.Operation)
		}
		latest[r.Operation] = r
	}

	summary := PermissionSummary{ApplicationID: applicationID, Total: len(order)}
	for _, op := range order {
		switch latest[op].State {
		case StateGranted, StateCompleted:
			summary.Granted = append(summary.Granted, op)
		case StateDenied, StateCancelled:
			summary.Denied = append(summary.Denied, op)
		case StateExpired:
			summary.Expired = append(summary.Expired, op)
		default:
			// pending-user-mediation and failed (recoverable) are unresolved.
			summary.Pending = append(summary.Pending, op)
		}
	}
	sort.Strings(summary.Granted)
	sort.Strings(summary.Pending)
	sort.Strings(summary.Denied)
	sort.Strings(summary.Expired)
	summary.AllResolved = len(summary.Pending) == 0
	return summary
}
