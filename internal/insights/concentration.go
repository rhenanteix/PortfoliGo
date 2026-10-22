package insights

import "fmt"

type ConcentrationRule struct {
	Threshold float64
}

func NewConcentrationRule(
	threshold float64,
) ConcentrationRule {

	return ConcentrationRule{
		Threshold: threshold,
	}
}

func (r ConcentrationRule) Evaluate(
	ctx Context,
) *Insight {

	concentration := ctx.Portfolio.Concentration

	if len(concentration) == 0 {
		return nil
	}

	largest := concentration[0]

	for _, payer := range concentration[1:] {

		if payer.Percentage >
			largest.Percentage {

			largest = payer
		}
	}

	if largest.Percentage <= r.Threshold {
		return nil
	}

	return &Insight{
		Code:      "HIGH_SINGLE_PAYER_CONCENTRATION",
		Title:     "Concentração por sacado",
		Severity:  SeverityWarning,
		Metric:    "single_payer_concentration",
		Value:     largest.Percentage,
		Threshold: r.Threshold,

		Message: fmt.Sprintf(
			"O maior sacado representa %.2f%% da carteira.",
			largest.Percentage,
		),

		Metadata: map[string]any{
			"payer_name":     largest.PayerName,
			"payer_document": largest.PayerDocument,
			"amount":         largest.Amount,
			"category":       CategoryPortfolio,
		},
	}
}
