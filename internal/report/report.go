package report

import "time"

// Report is the finished package the boss assembles and hands to the web layer
// (which serializes it to JSON for the React app).
//
// It deliberately has NO top-level success/verdict field: Stocker is research-
// only, so the report just presents the sections and lets each one tell its own
// truth (ok / flagged / failed). "Surface the tension, don't force an answer"
// lives in this shape.
type Report struct {
	Ticker      string    // which stock this report is about, e.g. "AAPL"
	Sections    []Section // the assembled worker outputs, each with its own Status
	GeneratedAt time.Time // when the boss assembled this report
}
