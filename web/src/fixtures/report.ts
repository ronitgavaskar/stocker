import type { Report } from "../types";

// The MIXED-STATE fixture: ownership failed (data null), froth flagged (data
// present + note), overview/financial ok. We build the UI against this on
// purpose so the failed/flagged/ok branching is structural, not bolted on after.
// This is a real payload the Go backend produced; it mirrors the wire shape 1:1.
export const mixedStateReport: Report = {
  ticker: "AAPL",
  generatedAt: "2026-06-07T14:23:57.910Z",
  sections: [
    {
      kind: "overview",
      status: "ok",
      note: "",
      data: {
        ticker: "AAPL",
        name: "Apple Inc.",
        summary:
          "Apple Inc. designs, manufactures, and markets smartphones, personal computers, tablets, wearables, and accessories, and sells related services.",
        sector: "TECHNOLOGY",
        industry: "ELECTRONIC COMPUTERS",
        marketCap: 3000000000000,
      },
    },
    {
      kind: "financial",
      status: "ok",
      note: "",
      data: {
        revenueTTM: 391035000000,
        grossProfitTTM: 180683000000,
        profitMargin: 0.247,
        returnOnEquityTTM: 1.5,
        eps: 6.13,
        ebitda: 134661000000,
      },
    },
    {
      kind: "ownership",
      status: "failed",
      note: "Alpha Vantage rate limit reached",
      data: null,
    },
    {
      kind: "froth",
      status: "flagged",
      note: "valuation looks stretched: trailing P/E 31.5 is well above the broad market",
      data: {
        beta: 1.25,
        pegRatio: 2.8,
        trailingPE: 31.5,
        forwardPE: 28.2,
        week52High: 260.1,
        week52Low: 164.08,
      },
    },
  ],
};
