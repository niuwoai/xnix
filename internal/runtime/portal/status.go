package portal

// UserStatus is a KDE-safe, user-facing projection of a portal request for
// tray, notification, and Compatibility Center surfaces. It exposes the safe
// operation and a friendly status plus the actions a user can take, never host
// paths, handle-token internals, or backend details.
type UserStatus struct {
	Operation             string   `json:"operation"`
	Status                string   `json:"status"`
	Actions               []string `json:"actions"`
	PermissionGranted     bool     `json:"permission_granted"`
	AwaitingUser          bool     `json:"awaiting_user"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

// UserStatus projects a request into a KDE-safe user-facing status.
func (r *Request) UserStatus() UserStatus {
	status := UserStatus{Operation: r.Operation}
	switch r.State {
	case StatePendingUserMediation:
		status.Status = "Waiting for your approval"
		status.Actions = []string{"approve", "deny"}
		status.AwaitingUser = true
	case StateGranted:
		status.Status = "Allowed"
		status.Actions = []string{"revoke"}
		status.PermissionGranted = true
	case StateCompleted:
		status.Status = "Allowed"
		status.PermissionGranted = true
	case StateDenied:
		status.Status = "Blocked"
		status.Actions = []string{"review-permissions"}
	case StateCancelled:
		status.Status = "Cancelled"
		status.Actions = []string{"request-again"}
	case StateFailed:
		status.Status = "Temporarily unavailable"
		status.Actions = []string{"retry"}
	case StateExpired:
		status.Status = "Request expired"
		status.Actions = []string{"request-again"}
	default:
		status.Status = "Unknown"
	}
	if status.Actions == nil {
		status.Actions = []string{}
	}
	return status
}
