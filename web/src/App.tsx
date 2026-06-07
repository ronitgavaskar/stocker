import { useState, useEffect, type FormEvent } from "react";
import type { Report } from "./types";
import { getReport } from "./api";
import { ReportView } from "./components/ReportView";
import "./App.css";

export default function App() {
  const [ticker, setTicker] = useState("AAPL");
  const [report, setReport] = useState<Report | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function load(t: string) {
    setLoading(true);
    setError(null);
    try {
      setReport(await getReport(t));
    } catch (err) {
      setError(err instanceof Error ? err.message : "Something went wrong");
      setReport(null);
    } finally {
      setLoading(false);
    }
  }

  // Load AAPL on first render so the styled report shows immediately.
  useEffect(() => {
    load("AAPL");
  }, []);

  function onSearch(e: FormEvent) {
    e.preventDefault();
    const t = ticker.trim();
    if (t) load(t);
  }

  return (
    <main className="app">
      <header className="app__brand">
        <h1>Stocker</h1>
        <p>Type a ticker, get a plain-English research one-pager.</p>
      </header>

      <form className="search" onSubmit={onSearch}>
        <input
          value={ticker}
          onChange={(e) => setTicker(e.target.value)}
          placeholder="e.g. AAPL"
          aria-label="Ticker"
        />
        <button type="submit" disabled={loading}>
          {loading ? "Researching…" : "Research"}
        </button>
      </form>

      {error && <p className="app__error">{error}</p>}
      {loading && <p className="app__status">Dispatching workers…</p>}
      {report && !loading && (
        <>
          <ReportView report={report} />
          <p className="app__disclaimer">
            Research only — not investment advice. Sections can disagree; we
            surface the tension rather than force a verdict.
          </p>
        </>
      )}
    </main>
  );
}
