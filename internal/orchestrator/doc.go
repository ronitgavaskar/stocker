// Package orchestrator is the boss. It fans a list of workers out
// concurrently and assembles their sections into one report.Report,
// surfacing disagreement rather than forcing a verdict. It defines the
// Worker interface it consumes, so it never imports package worker — main
// wires the concrete workers in.
package orchestrator
