package tool

import "context"

// Overview is the clean, internal "what is this company" shape returned by an
// overview tool. It is the domain type — real Go types, our own names — not the
// raw Alpha Vantage wire shape (which returns every value as a JSON string).
// Translating that messy wire shape into this is the anti-corruption layer's
// job; that raw struct lives with the real tool and is deferred for now.
//
// Each field notes its Alpha Vantage source key (function=OVERVIEW).
type Overview struct {
	Ticker    string `json:"ticker"`    // AV: Symbol
	Name      string `json:"name"`      // AV: Name
	Summary   string `json:"summary"`   // AV: Description
	Sector    string `json:"sector"`    // AV: Sector
	Industry  string `json:"industry"`  // AV: Industry
	MarketCap int64  `json:"marketCap"` // AV: MarketCapitalization (string -> int64), USD
}

// StubOverviewTool is a fake overview tool: it returns hardcoded AAPL data and
// never fails. It lets us build and run the orchestrator end-to-end with zero
// API calls (Alpha Vantage's free tier is harshly rate-limited). The real
// Alpha Vantage tool will satisfy the same method signature later.
type StubOverviewTool struct{}

// Fetch returns the overview for a ticker. The stub ignores ticker and always
// returns AAPL. The signature returns the error interface (not *ToolError) by
// design — the caller extracts the kind via errors.As when it's non-nil.
func (StubOverviewTool) Fetch(ctx context.Context, ticker string) (Overview, error) {
	return Overview{
		Ticker:    "AAPL",
		Name:      "Apple Inc.",
		Summary:   "Apple Inc. designs, manufactures, and markets smartphones, personal computers, tablets, wearables, and accessories, and sells related services.",
		Sector:    "TECHNOLOGY",         // AV returns these uppercase; value-normalization is a later choice
		Industry:  "ELECTRONIC COMPUTERS",
		MarketCap: 3_000_000_000_000, // ~$3T, stub value
	}, nil
}
