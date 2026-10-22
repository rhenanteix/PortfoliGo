package analysis

import (
	"testing"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func TestCalculateConcentration(t *testing.T) {

	items := []receivable.Receivable{
		{
			PayerDocument: "111",
			PayerName:     "Empresa A",
			Amount:        700000,
		},
		{
			PayerDocument: "222",
			PayerName:     "Empresa B",
			Amount:        300000,
		},
	}

	result := CalculateConcentration(items)

	if len(result) != 2 {
		t.Fatalf(
			"expected 2 payers, got %d",
			len(result),
		)
	}

	if result[0].PayerDocument != "111" {
		t.Errorf(
			"expected payer 111 first",
		)
	}

	if result[0].Percentage != 70 {
		t.Errorf(
			"expected 70%%, got %.2f",
			result[0].Percentage,
		)
	}
}
