package worker

import (
	"context"
	"time"

	"github.com/ronitgavaskar/stocker/internal/report"
	"github.com/ronitgavaskar/stocker/internal/tool"
)

const (
	financialKind   = "financial"
	financialBudget = 12 * time.Second
)

// FinancialTool is what the financial worker needs from a tool. Defined here
// (the consumer) so the stub and the real Alpha Vantage tool satisfy it
// implicitly.
type FinancialTool interface {
	Fetch(ctx context.Context, ticker string) (tool.Financial, error)
}

// FinancialWorker produces the "financial health" section. Satisfies
// orchestrator.Worker implicitly via Run; ticker is bound at construction.
type FinancialWorker struct {
	tool   FinancialTool
	ticker string
}

// NewFinancial injects the tool and binds the ticker for this request.
func NewFinancial(t FinancialTool, ticker string) *FinancialWorker {
	return &FinancialWorker{tool: t, ticker: ticker}
}

// Run applies the 12s budget, calls the tool, and always returns a Section.
func (w *FinancialWorker) Run(ctx context.Context) report.Section {
	ctx, cancel := context.WithTimeout(ctx, financialBudget)
	defer cancel()

	data, err := w.tool.Fetch(ctx, w.ticker)
	if err != nil {
		return failedSection(financialKind, err)
	}
	return report.Section{Kind: financialKind, Status: report.StatusOK, Data: data}
}
