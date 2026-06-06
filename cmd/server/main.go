// Command server is Stocker's entrypoint. For now it runs a one-shot smoke test
// of the concurrent pipe (stub -> worker -> orchestrator -> report); it becomes
// the HTTP server in the next step.
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ronitgavaskar/stocker/internal/orchestrator"
	"github.com/ronitgavaskar/stocker/internal/tool"
	"github.com/ronitgavaskar/stocker/internal/worker"
)

func main() {
	// Shared tool: constructed once, reused across requests and across workers.
	stub := tool.StubOverviewTool{}

	// The worker factory: main is the only place that knows the concrete
	// workers. Called fresh per request so each worker carries that ticker.
	makeWorkers := func(ticker string) []orchestrator.Worker {
		return []orchestrator.Worker{
			worker.NewOverview(stub, ticker),
		}
	}

	boss := orchestrator.New(makeWorkers)

	// Run the concurrent pipe once and print the assembled report.
	rep := boss.Assemble(context.Background(), "AAPL")

	fmt.Printf("Report: %s   (generated %s)\n", rep.Ticker, rep.GeneratedAt.Format(time.RFC3339))
	for _, s := range rep.Sections {
		line := fmt.Sprintf("  %-10s %s", s.Kind, s.Status)
		if s.Note != "" {
			line += "  — " + s.Note
		}
		fmt.Println(line)
	}
}
