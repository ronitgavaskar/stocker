package tool

import "context"

// Ownership is the clean, internal "who holds it and how do analysts feel"
// shape. Percentage fields are FRACTIONS (e.g. 0.247 = 24.7%) — one rule
// everywhere.
//
// Seam note for the real tool: AV's OVERVIEW returns PercentInstitutions as a
// percent NUMBER (e.g. 65.626), not a fraction, so the real tool must divide
// the Percent* fields by 100 to honor this convention. The stub uses fractions
// directly.
type Ownership struct {
	PercentInsiders     float64        `json:"percentInsiders"`     // fraction, e.g. 0.247 = 24.7%
	PercentInstitutions float64        `json:"percentInstitutions"` // fraction, e.g. 0.247 = 24.7%
	AnalystTargetPrice  float64        `json:"analystTargetPrice"`  // AV: AnalystTargetPrice (USD)
	Ratings             AnalystRatings `json:"ratings"`
}

// AnalystRatings groups AV's five AnalystRating* count fields.
type AnalystRatings struct {
	StrongBuy  int `json:"strongBuy"`  // AV: AnalystRatingStrongBuy
	Buy        int `json:"buy"`        // AV: AnalystRatingBuy
	Hold       int `json:"hold"`       // AV: AnalystRatingHold
	Sell       int `json:"sell"`       // AV: AnalystRatingSell
	StrongSell int `json:"strongSell"` // AV: AnalystRatingStrongSell
}

// StubOwnershipTool returns hardcoded AAPL ownership/sentiment and never fails.
type StubOwnershipTool struct{}

// Fetch ignores ticker and always returns AAPL.
func (StubOwnershipTool) Fetch(ctx context.Context, ticker string) (Ownership, error) {
	return Ownership{
		PercentInsiders:     0.0006, // ~0.06%
		PercentInstitutions: 0.62,   // ~62%
		AnalystTargetPrice:  250.50,
		Ratings: AnalystRatings{
			StrongBuy:  28,
			Buy:        21,
			Hold:       14,
			Sell:       2,
			StrongSell: 1,
		},
	}, nil
}
