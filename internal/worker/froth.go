package worker

import (
	"context"
	"time"

	"github.com/ronitgavaskar/stocker/internal/report"
	"github.com/ronitgavaskar/stocker/internal/tool"
)

const (
	frothKind   = "froth"
	frothBudget = 5 * time.Second
)

// FrothTool is what the froth worker needs from a tool. Defined here (the
// consumer) so the stub and the real Alpha Vantage tool satisfy it implicitly.
type FrothTool interface {
	Fetch(ctx context.Context, ticker string) (tool.Froth, error)
}

// FrothWorker produces the "froth check" section. Satisfies orchestrator.Worker
// implicitly via Run; ticker is bound at construction.
type FrothWorker struct {
	tool   FrothTool
	ticker string
}

// NewFroth injects the tool and binds the ticker for this request.
func NewFroth(t FrothTool, ticker string) *FrothWorker {
	return &FrothWorker{tool: t, ticker: ticker}
}

// Run applies the 5s budget, calls the tool, and always returns a Section.
func (w *FrothWorker) Run(ctx context.Context) report.Section {
	ctx, cancel := context.WithTimeout(ctx, frothBudget)
	defer cancel()

	data, err := w.tool.Fetch(ctx, w.ticker)
	if err != nil {
		return failedSection(frothKind, err)
	}
	return report.Section{Kind: frothKind, Status: report.StatusOK, Data: data}
}
