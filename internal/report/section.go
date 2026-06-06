package report

// Status is the outcome of a worker producing a Section. Its zero value is
// StatusUnknown, so a Section nobody filled in never masquerades as success.
type Status int

const (
	StatusUnknown Status = iota // 0 — unset/invalid; the safe default
	StatusOK                    // full, trustworthy content
	StatusFlagged               // content present, but with a caveat
	StatusFailed                // no usable content; Note says why
)

// String renders a Status as a human-readable word for logs and demo output.
// This is the DISPLAY stringer; the JSON-wire shape of Status is a separate
// decision still deferred to the httpapi step.
func (s Status) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusFlagged:
		return "Flagged"
	case StatusFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}

// Section is one worker's contribution to the report. A worker ALWAYS returns
// a Section, never an error: a failure is just a Section with a failed/flagged
// Status and a Note. Content (a Data field) is added in a later step.
type Section struct {
	Kind   string // which section: "overview", "financial", ...
	Status Status
	Note   string // failure reason or caveat; empty when OK

	// TODO(data): add a `Data any` field to carry the section's content once
	// a worker produces real fields. Deferred on purpose — adding an exported
	// field later is backward-compatible, so nothing breaks by waiting.
}
