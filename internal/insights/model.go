package insights

type Severity string

const (
	SeverityInfo     Severity = "info"
	SeverityWarning  Severity = "warning"
	SeverityCritical Severity = "critical"
)

type Category string

const (
	CategoryPortfolio   Category = "portfolio"
	CategoryMaturity    Category = "maturity"
	CategoryDataQuality Category = "data_quality"
)

type Insight struct {
	Code      string         `json:"code"`
	Title     string         `json:"title"`
	Message   string         `json:"message"`
	Severity  Severity       `json:"severity"`
	Category  Category       `json:"category"`
	Metric    string         `json:"metric,omitempty"`
	Value     float64        `json:"value,omitempty"`
	Threshold float64        `json:"threshold,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}