package orchestrator

import (
	"context"

	"github.com/ronitgavaskar/stocker/internal/report"
)

// Worker is one unit of research the boss dispatches. The boss CONSUMES this
// interface, so it lives here in package orchestrator, not in package worker.
// Concrete workers in package worker satisfy it implicitly — Go has no
// `implements` keyword — which is why orchestrator and worker never import
// each other; only cmd/server wires the concrete workers in.
//
// Run takes a context for cancellation and deadlines. The worker tightens that
// context to its own time budget internally (e.g. overview 20s) and owns its
// retry policy.
//
// Run ALWAYS returns a report.Section and never an error: a failure is just a
// Section with a failed/flagged Status and a Note. That lets the boss treat
// every worker outcome through one uniform path — collect the Section, move on.
type Worker interface {
	Run(ctx context.Context) report.Section
}
