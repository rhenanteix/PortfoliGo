package validation

import (
	"time"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

type DueDateValidator struct {
	Now time.Time
}

func (v DueDateValidator) Validate(
	item receivable.Receivable,
	row int,
) []Finding {

	today := time.Date(
		v.Now.Year(),
		v.Now.Month(),
		v.Now.Day(),
		0,
		0,
		0,
		0,
		v.Now.Location(),
	)

	if item.DueDate.Before(today) {
		return []Finding{
			{
				Row:      row,
				Code:     "OVERDUE_RECEIVABLE",
				Severity: SeverityWarning,
				Field:    "due_date",
				Message:  "receivable is overdue",
			},
		}
	}

	return nil
}
