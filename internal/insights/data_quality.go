package insights

import "fmt"

type DataQualityRule struct {
	Threshold float64
}

func NewDataQualityRule(
	threshold float64,
) DataQualityRule {

	return DataQualityRule{
		Threshold: threshold,
	}
}

func (r DataQualityRule) Evaluate(
	ctx Context,
) *Insight {

	total := ctx.DataQuality.Total

	if total == 0 {
		return nil
	}

	problematic :=
		ctx.DataQuality.Invalid +
			ctx.DataQuality.Review

	percentage :=
		float64(problematic) /
			float64(total) *
			100

	if percentage <= r.Threshold {
		return nil
	}

	return &Insight{
		Code:      "DATA_QUALITY_REVIEW",
		Title:     "Qualidade dos dados",
		Severity:  SeverityWarning,
		Metric:    "problematic_rows_percentage",
		Value:     percentage,
		Threshold: r.Threshold,

		Message: fmt.Sprintf(
			"%.2f%% dos registros requerem revisão.",
			percentage,
		),

		Metadata: map[string]any{
			"total":   total,
			"review":  ctx.DataQuality.Review,
			"invalid": ctx.DataQuality.Invalid,
			"category": CategoryDataQuality,
		},
	}
}
