package portal

import "testing"

func TestSummarizeBucketsLatestStatePerOperation(t *testing.T) {
	broker := NewFakeBroker()
	app := "org.example.ledger"

	// file-open: granted.
	fileReq, _ := broker.CreateRequest(RequestSpec{ApplicationID: app, Operation: "file-open"})
	broker.Resolve(fileReq.HandleToken, OutcomeGranted)
	// print: pending.
	broker.CreateRequest(RequestSpec{ApplicationID: app, Operation: "print"})
	// clipboard: expired.
	clip, _ := broker.CreateRequest(RequestSpec{ApplicationID: app, Operation: "clipboard"})
	broker.Expire(clip.HandleToken)
	// camera: denied by policy at creation.
	broker.CreateRequest(RequestSpec{ApplicationID: app, Operation: "camera"})
	// A different app's request must be excluded.
	broker.CreateRequest(RequestSpec{ApplicationID: "org.other.app", Operation: "file-open"})

	summary := Summarize(broker.List(), app)
	if summary.ApplicationID != app || summary.Total != 4 {
		t.Fatalf("unexpected summary totals: %#v", summary)
	}
	if !equalStrings(summary.Granted, []string{"file-open"}) ||
		!equalStrings(summary.Pending, []string{"print"}) ||
		!equalStrings(summary.Expired, []string{"clipboard"}) ||
		!equalStrings(summary.Denied, []string{"camera"}) {
		t.Fatalf("unexpected buckets: %#v", summary)
	}
	if summary.AllResolved {
		t.Fatalf("a pending request must make the summary unresolved")
	}
	if summary.BackendDetailsExposed {
		t.Fatalf("summary must not expose backend details")
	}
}

func TestSummarizeLatestRequestWinsPerOperation(t *testing.T) {
	broker := NewFakeBroker()
	app := "org.example.ledger"

	// First file-open fails, then a retry is granted: latest (granted) wins.
	first, _ := broker.CreateRequest(RequestSpec{ApplicationID: app, Operation: "file-open"})
	broker.Resolve(first.HandleToken, OutcomeFailed)
	second, _ := broker.CreateRequest(RequestSpec{ApplicationID: app, Operation: "file-open"})
	broker.Resolve(second.HandleToken, OutcomeGranted)

	summary := Summarize(broker.List(), app)
	if summary.Total != 1 || !equalStrings(summary.Granted, []string{"file-open"}) || len(summary.Pending) != 0 {
		t.Fatalf("latest request should win: %#v", summary)
	}
	if !summary.AllResolved {
		t.Fatalf("all resolved when the latest is granted")
	}
}

func TestSummarizeEmpty(t *testing.T) {
	summary := Summarize(nil, "org.example.ledger")
	if summary.Total != 0 || !summary.AllResolved || len(summary.Granted) != 0 {
		t.Fatalf("empty summary wrong: %#v", summary)
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
