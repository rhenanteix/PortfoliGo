package analysis

type PortfolioSummary struct {
	TotalAmount             int64 `json:"total_amount"`
	Receivables             int   `json:"receivables"`
	UniquePayers            int   `json:"unique_payers"`
	AverageTicket           int64 `json:"average_ticket"`
	WeightedAverageTermDays int   `json:"weighted_average_term_days"`
}

type PayerConcentration struct {
	PayerDocument string  `json:"payer_document"`
	PayerName     string  `json:"payer_name"`
	Amount        int64   `json:"amount"`
	Percentage    float64 `json:"percentage"`
}

type Aging struct {
	Overdue    int64 `json:"overdue"`
	Days0To30  int64 `json:"days_0_30"`
	Days31To60 int64 `json:"days_31_60"`
	Days61To90 int64 `json:"days_61_90"`
	Over90Days int64 `json:"over_90_days"`
}

type PortfolioAnalysis struct {
	Summary       PortfolioSummary     `json:"summary"`
	Concentration []PayerConcentration `json:"concentration"`
	Aging         Aging                `json:"aging"`
}
