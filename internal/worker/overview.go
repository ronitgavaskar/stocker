package worker

import (
	"context"
	"errors"
	"time"

	"github.com/ronitgavaskar/stocker/internal/report"
	"github.com/ronitgavaskar/stocker/internal/tool"
)

const (
	overviewKind   = "overview"
	overviewBudget = 20 * time.Second
)

// OverviewTool is what the overview worker needs from a tool. It is defined
// here — in the consumer — so that tool.StubOverviewTool, and later the real
// Alpha Vantage tool, satisfy it implicitly (Go has no `implements`). The
// worker imports package tool only for the Overview return type, never the
// concrete tool, which keeps it swappable and testable with a fake.
type OverviewTool interface {
	Fetch(ctx context.Context, ticker string) (tool.Overview, error)
}

// OverviewWorker produces the "what is this company" section. It satisfies
// orchestrator.Worker implicitly via Run; it does not import orchestrator.
//
// The ticker is bound at construction, not passed to Run: the Worker interface
// is Run(ctx) report.Section with no ticker, so workers are cheap per-request
// values that carry their own input. That keeps Run's signature uniform across
// every worker.
type OverviewWorker struct {
	tool   OverviewTool
	ticker string
}

// NewOverview injects the tool dependency and binds the ticker for this request.
func NewOverview(t OverviewTool, ticker string) *OverviewWorker {
	return &OverviewWorker{tool: t, ticker: ticker}
}

// Run applies the worker's own time budget, calls the tool, and ALWAYS returns
// a report.Section — success or failure. It never returns an error: a failure
// is a Section with StatusFailed and a Note. (Retries are a later step.)
func (w *OverviewWorker) Run(ctx context.Context) report.Section {
	ctx, cancel := context.WithTimeout(ctx, overviewBudget)
	defer cancel()

	ov, err := w.tool.Fetch(ctx, w.ticker)
	if err != nil {
		return failedSection(overviewKind, err)
	}

	_ = ov // TODO(data): carry overview content once Section.Data exists
	return report.Section{Kind: overviewKind, Status: report.StatusOK}
}

// failedSection turns any tool error into a failed Section instead of crashing —
// the result-with-status payoff. It digs out the *tool.ToolError via errors.As
// for a clean message, falling back to the raw error string otherwise.
func failedSection(kind string, err error) report.Section {
	note := err.Error()
	var te *tool.ToolError
	if errors.As(err, &te) {
		note = te.Message
	}
	return report.Section{Kind: kind, Status: report.StatusFailed, Note: note}
}
