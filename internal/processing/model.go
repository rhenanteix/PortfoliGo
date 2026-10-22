package processing

import (
	"github.com/rhenanteix/goldscan/internal/receivable"
	"github.com/rhenanteix/goldscan/internal/validation"
)

type ProcessedRow struct {
	Row      int                    `json:"row"`
	Raw      receivable.Input       `json:"raw"`
	Item     *receivable.Receivable `json:"item,omitempty"`
	Findings []validation.Finding   `json:"findings,omitempty"`
}
