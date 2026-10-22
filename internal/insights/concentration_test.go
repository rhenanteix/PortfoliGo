package insights

import (
	"testing"

	"github.com/rhenanteix/goldscan/internal/analysis"
)

func TestConcentrationRuleGeneratesInsight(
	t *testing.T,
) {

	ctx := Context{
		Portfolio: analysis.PortfolioAnalysis{
			Concentration: []analysis.PayerConcentration{
				{
					PayerName:     "Empresa Alpha",
					PayerDocument: "123",
					Amount:        400000,
					Percentage:    40,
				},
				{
					PayerName:     "Empresa Beta",
					PayerDocument: "456",
					Amount:        600000,
					Percentage:    60,
				},
			},
		},
	}

	rule := NewConcentrationRule(30)

	result := rule.Evaluate(ctx)

	if result == nil {
		t.Fatal(
			"expected insight, got nil",
		)
	}

	if result.Code !=
		"HIGH_SINGLE_PAYER_CONCENTRATION" {

		t.Errorf(
			"unexpected code: %s",
			result.Code,
		)
	}
}
