package validation

import (
	"fmt"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func FindDuplicates(
	items []receivable.Receivable,
) []Finding {

	seen := make(map[string]int)

	var findings []Finding

	for index, item := range items {

		row := index + 2

		key := duplicateKey(item)

		firstRow, exists := seen[key]

		if exists {
			findings = append(
				findings,
				Finding{
					Row:      row,
					Code:     "POSSIBLE_DUPLICATE",
					Severity: SeverityWarning,
					Message: fmt.Sprintf(
						"possible duplicate of row %d",
						firstRow,
					),
				},
			)

			continue
		}

		seen[key] = row
	}

	return findings
}

func duplicateKey(
	item receivable.Receivable,
) string {

	return fmt.Sprintf(
		"%s|%s|%d|%s",
		item.PayerDocument,
		item.InvoiceNumber,
		item.Amount,
		item.DueDate.Format("2006-01-02"),
	)
}
