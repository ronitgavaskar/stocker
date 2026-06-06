// Package worker holds the concurrent worker agents. Each worker owns one
// report section, calls a tool, and owns the how-much of retrying (backoff,
// attempt count, remaining time budget). Workers satisfy the
// orchestrator.Worker interface implicitly, so this package does not import
// orchestrator.
package worker
