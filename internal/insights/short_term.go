package insights

import "fmt"

type ShortTermRule struct {
	Threshold float64
}

func NewShortTermRule(
	threshold float64,
) ShortTermRule {

	return ShortTermRule{
		Threshold: threshold,
	}
}

func (r ShortTermRule) Evaluate(
	ctx Context,
) *Insight {

	total := ctx.Portfolio.Summary.TotalAmount

	if total <= 0 {
		return nil
	}

	shortTerm :=
		ctx.Portfolio.Aging.Days0To30

	percentage :=
		float64(shortTerm) /
			float64(total) *
			100

	if percentage <= r.Threshold {
		return nil
	}

	return &Insight{
		Code:      "SHORT_TERM_CONCENTRATION",
		Title:     "Concentração de vencimentos",
		Severity:  SeverityInfo,
		Metric:    "days_0_30_percentage",
		Value:     percentage,
		Threshold: r.Threshold,

		Message: fmt.Sprintf(
			"%.2f%% do valor da carteira vence nos próximos 30 dias.",
			percentage,
		),

		Metadata: map[string]any{
			"amount": shortTerm,
			"category": CategoryMaturity,
		},
	}
}
