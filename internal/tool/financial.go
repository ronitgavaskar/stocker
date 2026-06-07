package tool

import "context"

// Financial is the clean, internal "is this a healthy, profitable business"
// shape. Each field notes its Alpha Vantage source key (function=OVERVIEW).
// Percentage fields are FRACTIONS (e.g. 0.247 = 24.7%) — one rule everywhere.
type Financial struct {
	RevenueTTM        int64   `json:"revenueTTM"`        // AV: RevenueTTM (USD)
	GrossProfitTTM    int64   `json:"grossProfitTTM"`    // AV: GrossProfitTTM (USD)
	ProfitMargin      float64 `json:"profitMargin"`      // fraction, e.g. 0.247 = 24.7%
	ReturnOnEquityTTM float64 `json:"returnOnEquityTTM"` // fraction, e.g. 0.247 = 24.7%
	EPS               float64 `json:"eps"`               // AV: EPS
	EBITDA            int64   `json:"ebitda"`            // AV: EBITDA (USD)
}

// StubFinancialTool returns hardcoded AAPL financials and never fails.
type StubFinancialTool struct{}

// Fetch ignores ticker and always returns AAPL. Signature returns the error
// interface (not *ToolError) by design; the caller uses errors.As when non-nil.
func (StubFinancialTool) Fetch(ctx context.Context, ticker string) (Financial, error) {
	return Financial{
		RevenueTTM:        391_035_000_000, // ~$391B
		GrossProfitTTM:    180_683_000_000, // ~$181B
		ProfitMargin:      0.247,
		ReturnOnEquityTTM: 1.50, // ~150%; AAPL's equity is shrunk by buybacks
		EPS:               6.13,
		EBITDA:            134_661_000_000, // ~$135B
	}, nil
}
