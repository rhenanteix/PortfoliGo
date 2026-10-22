package mapping

type Status string

const (
	StatusReady           Status = "ready"
	StatusMappingRequired Status = "mapping_required"
)

type ColumnSuggestion struct {
	SourceColumn string  `json:"source_column"`
	Field        Field   `json:"field,omitempty"`
	Recognized   bool    `json:"recognized"`
	Confidence float64 `json:"confidence"`	
}

type Result struct {
	Status             Status             `json:"status"`
	Columns            []ColumnSuggestion `json:"columns"`
	Mapping            ColumnMapping      `json:"mapping"`
	Unrecognized       []string           `json:"unrecognized"`
	MissingRequired    []Field            `json:"missing_required"`
	RecognitionPercent float64            `json:"recognition_percent"`
}