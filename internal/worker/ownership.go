package worker

import (
	"context"
	"time"

	"github.com/ronitgavaskar/stocker/internal/report"
	"github.com/ronitgavaskar/stocker/internal/tool"
)

const (
	ownershipKind = "ownership"
	// ownershipBudget: no budget was locked for ownership (only overview 20s /
	// financial 12s / froth 5s). 10s is a placeholder — revisit with retries.
	ownershipBudget = 10 * time.Second
)

// OwnershipTool is what the ownership worker needs from a tool. Defined here
// (the consumer) so the stub and the real Alpha Vantage tool satisfy it
// implicitly.
type OwnershipTool interface {
	Fetch(ctx context.Context, ticker string) (tool.Ownership, error)
}

// OwnershipWorker produces the "ownership / sentiment" section. Satisfies
// orchestrator.Worker implicitly via Run; ticker is bound at construction.
type OwnershipWorker struct {
	tool   OwnershipTool
	ticker string
}

// NewOwnership injects the tool and binds the ticker for this request.
func NewOwnership(t OwnershipTool, ticker string) *OwnershipWorker {
	return &OwnershipWorker{tool: t, ticker: ticker}
}

// Run applies the budget, calls the tool, and always returns a Section.
func (w *OwnershipWorker) Run(ctx context.Context) report.Section {
	ctx, cancel := context.WithTimeout(ctx, ownershipBudget)
	defer cancel()

	data, err := w.tool.Fetch(ctx, w.ticker)
	if err != nil {
		return failedSection(ownershipKind, err)
	}
	return report.Section{Kind: ownershipKind, Status: report.StatusOK, Data: data}
}
