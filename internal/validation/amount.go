package validation

import "github.com/rhenanteix/goldscan/internal/receivable"

type AmountValidator struct{}

func (AmountValidator) Validate(
	item receivable.Receivable,
	row int,
) []Finding {

	if item.Amount <= 0 {
		return []Finding{
			{
				Row:      row,
				Code:     "INVALID_AMOUNT",
				Severity: SeverityError,
				Field:    "amount",
				Message:  "amount must be greater than zero",
			},
		}
	}

	return nil
}
