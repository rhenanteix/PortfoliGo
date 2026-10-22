package validation

type Severity string

const (
	SeverityInfo    Severity = "info"
	SeverityWarning Severity = "warning"
	SeverityError   Severity = "error"
)

type Finding struct {
	Row      int      `json:"row"`
	Code     string   `json:"code"`
	Severity Severity `json:"severity"`
	Field    string   `json:"field,omitempty"`
	Message  string   `json:"message"`
}

type Summary struct {
	Total   int `json:"total"`
	Valid   int `json:"valid"`
	Review  int `json:"review"`
	Invalid int `json:"invalid"`
}
