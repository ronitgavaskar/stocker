package orchestrator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ronitgavaskar/stocker/internal/report"
)

// WorkerFactory builds the roster of workers for a single request. main owns
// the implementation — it is the one place allowed to import package worker —
// so the orchestrator only ever holds this function value and stays blind to
// which concrete workers exist. It is called fresh per request because workers
// carry that request's ticker.
type WorkerFactory func(ticker string) []Worker

// Orchestrator is the boss: a thin, worker-agnostic dispatcher. It runs a
// roster of workers concurrently and assembles their Sections into one Report.
// It forces no verdict, and a single misbehaving worker can never take down the
// report or the process.
type Orchestrator struct {
	makeWorkers WorkerFactory
}

// New injects the worker factory at startup.
func New(makeWorkers WorkerFactory) *Orchestrator {
	return &Orchestrator{makeWorkers: makeWorkers}
}

// Assemble builds the workers for this ticker, runs them all concurrently, and
// gathers their Sections into a Report. Every worker outcome — success, clean
// failure, or even a panic — becomes a Section, so one bad worker costs exactly
// one section, never the whole report.
func (o *Orchestrator) Assemble(ctx context.Context, ticker string) report.Report {
	workers := o.makeWorkers(ticker)

	// One result slot per worker. Each goroutine writes ONLY to its own index,
	// so although the slice is shared, two goroutines never touch the same
	// element — no mutex needed. Slot order = dispatch order = report order.
	sections := make([]report.Section, len(workers))

	var wg sync.WaitGroup
	wg.Add(len(workers))
	for i, w := range workers {
		// Go 1.22+: i and w are fresh each iteration, so capturing them in the
		// closure is safe (the classic loop-variable capture bug is gone).
		go func() {
			// defer order is LIFO and deliberate: wg.Done is registered FIRST
			// so it runs LAST — after the recover below has written the section.
			// Otherwise Wait() could return before a panicked slot is filled.
			defer wg.Done()

			// recover() firewall: an unrecovered panic in ANY goroutine crashes
			// the whole process. Catch it here and turn it into a failed Section
			// so a freak worker bug costs one section, not the server.
			defer func() {
				if r := recover(); r != nil {
					sections[i] = report.Section{
						Kind:   "unknown", // Worker has no Kind() yet (Q2=A); can't name the panicker
						Status: report.StatusFailed,
						Note:   fmt.Sprintf("worker panicked: %v", r),
					}
				}
			}()

			sections[i] = w.Run(ctx)
		}()
	}
	wg.Wait()

	return report.Report{
		Ticker:      ticker,
		Sections:    sections,
		GeneratedAt: time.Now().UTC(),
	}
}
