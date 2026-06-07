package report

import "encoding/json"

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
// This is the DISPLAY form, kept deliberately independent of the JSON wire form
// (see MarshalJSON): log text can be reworded without breaking the API contract.
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

// MarshalJSON renders a Status as a lowercase string token on the wire. This is
// an INDEPENDENT switch, deliberately NOT derived from String(): the JSON is an
// API contract the frontend depends on, whereas String() is human-facing log
// text that must be free to change without silently altering that contract.
func (s Status) MarshalJSON() ([]byte, error) {
	var token string
	switch s {
	case StatusOK:
		token = "ok"
	case StatusFlagged:
		token = "flagged"
	case StatusFailed:
		token = "failed"
	default:
		token = "unknown"
	}
	return json.Marshal(token)
}

// Section is one worker's contribution to the report. A worker ALWAYS returns
// a Section, never an error: a failure is just a Section with a failed/flagged
// Status and a Note.
//
// No field uses omitempty: kind/status/note/data are always present in the JSON,
// so the frontend never branches on whether a field exists. On success Data
// holds the worker's content struct (a VALUE type, e.g. tool.Overview); on
// failure Data is simply left unset, so it stays nil and marshals to null — not
// {}. Using value types (not pointers) in Data also avoids the typed-nil-in-
// interface trap, where a non-nil interface wrapping a nil pointer surprises you.
type Section struct {
	Kind   string `json:"kind"`   // which section: "overview", "financial", ...
	Status Status `json:"status"` // lowercase string on the wire (see MarshalJSON)
	Note   string `json:"note"`   // failure reason or caveat; empty but always present
	Data   any    `json:"data"`   // content struct on success; nil -> null on failure
}
