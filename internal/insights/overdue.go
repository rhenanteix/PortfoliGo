package insights

import "fmt"

type OverdueRule struct {
	Threshold float64
}

func NewOverdueRule(
	threshold float64,
) OverdueRule {

	return OverdueRule{
		Threshold: threshold,
	}
}

func (r OverdueRule) Evaluate(
	ctx Context,
) *Insight {

	total := ctx.Portfolio.Summary.TotalAmount

	if total <= 0 {
		return nil
	}

	overdue := ctx.Portfolio.Aging.Overdue

	percentage :=
		float64(overdue) /
			float64(total) *
			100

	if percentage <= r.Threshold {
		return nil
	}

	return &Insight{
		Code:      "OVERDUE_EXPOSURE",
		Title:     "Títulos vencidos",
		Severity:  SeverityWarning,
		Metric:    "overdue_percentage",
		Value:     percentage,
		Threshold: r.Threshold,

		Message: fmt.Sprintf(
			"%.2f%% do valor processável da carteira está vencido.",
			percentage,
		),

		Metadata: map[string]any{
			"amount": overdue,
			"category": CategoryMaturity,
		},
	}
}
