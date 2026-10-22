package validation

import (
	"testing"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func TestAmountValidator(t *testing.T) {

	validator := AmountValidator{}

	tests := []struct {
		name             string
		amount           int64
		expectedFindings int
	}{
		{
			name:             "valid amount",
			amount:           10000,
			expectedFindings: 0,
		},
		{
			name:             "zero amount",
			amount:           0,
			expectedFindings: 1,
		},
		{
			name:             "negative amount",
			amount:           -100,
			expectedFindings: 1,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			item := receivable.Receivable{
				Amount: test.amount,
			}

			findings := validator.Validate(
				item,
				2,
			)

			if len(findings) != test.expectedFindings {
				t.Errorf(
					"expected %d findings, got %d",
					test.expectedFindings,
					len(findings),
				)
			}
		})
	}
}
