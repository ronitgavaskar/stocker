// Command server is Stocker's entrypoint: it wires the tool, workers,
// orchestrator, and HTTP API together and serves the JSON API.
package main

import (
	"log"
	"net/http"

	"github.com/ronitgavaskar/stocker/internal/httpapi"
	"github.com/ronitgavaskar/stocker/internal/orchestrator"
	"github.com/ronitgavaskar/stocker/internal/tool"
	"github.com/ronitgavaskar/stocker/internal/worker"
)

func main() {
	// Shared tools: constructed once, reused across requests and across workers.
	overviewTool := tool.StubOverviewTool{}
	financialTool := tool.StubFinancialTool{}
	ownershipTool := tool.StubOwnershipTool{}
	frothTool := tool.StubFrothTool{}

	// The worker factory: main is the only place that knows the concrete
	// workers. Called fresh per request so each worker carries that ticker.
	// Slice order = report section order (overview, financial, ownership, froth).
	makeWorkers := func(ticker string) []orchestrator.Worker {
		return []orchestrator.Worker{
			worker.NewOverview(overviewTool, ticker),
			worker.NewFinancial(financialTool, ticker),
			worker.NewOwnership(ownershipTool, ticker),
			worker.NewFroth(frothTool, ticker),
		}
	}

	boss := orchestrator.New(makeWorkers)
	api := httpapi.NewHandler(boss)

	const addr = ":8080"
	log.Printf("stocker listening on %s", addr)
	if err := http.ListenAndServe(addr, api.Routes()); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
