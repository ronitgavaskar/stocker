package tool

import "context"

// Froth is the clean, internal "hype-risk / valuation-stretch signals" shape.
// It only CARRIES the numbers; deciding "this looks frothy -> flag it" is worker
// logic for later. Each field notes its Alpha Vantage source key.
type Froth struct {
	Beta       float64 `json:"beta"`       // AV: Beta (volatility vs market)
	PEGRatio   float64 `json:"pegRatio"`   // AV: PEGRatio (growth-adjusted valuation)
	TrailingPE float64 `json:"trailingPE"` // AV: TrailingPE
	ForwardPE  float64 `json:"forwardPE"`  // AV: ForwardPE (trailing vs forward = expectations gap)
	Week52High float64 `json:"week52High"` // AV: 52WeekHigh (USD)
	Week52Low  float64 `json:"week52Low"`  // AV: 52WeekLow (USD)
}

// StubFrothTool returns hardcoded AAPL froth signals and never fails.
type StubFrothTool struct{}

// Fetch ignores ticker and always returns AAPL.
func (StubFrothTool) Fetch(ctx context.Context, ticker string) (Froth, error) {
	return Froth{
		Beta:       1.25,
		PEGRatio:   2.80,
		TrailingPE: 31.50,
		ForwardPE:  28.20,
		Week52High: 260.10,
		Week52Low:  164.08,
	}, nil
}
