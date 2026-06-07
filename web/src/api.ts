import type { Report } from "./types";
import { mixedStateReport } from "./fixtures/report";

// getReport currently returns the hardcoded MIXED-STATE fixture so the UI is
// built against the worst case (failed/flagged/ok) from line one. Swapping this
// for a real `fetch("/report?ticker=" + ticker)` is a later step — and that is
// when the CORS / Vite-proxy decision comes due.
export async function getReport(ticker: string): Promise<Report> {
  // Simulate latency so the loading state is exercised during development.
  await new Promise((resolve) => setTimeout(resolve, 300));
  return { ...mixedStateReport, ticker: ticker.toUpperCase() };
}
