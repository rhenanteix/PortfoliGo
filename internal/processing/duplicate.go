package processing

import (
	"fmt"

	"github.com/rhenanteix/goldscan/internal/receivable"
	"github.com/rhenanteix/goldscan/internal/validation"
)

func FindDuplicates(
	rows []ProcessedRow,
) []validation.Finding {

	seen := make(map[string]int)

	var findings []validation.Finding

	for _, row := range rows {

		if row.Item == nil {
			continue
		}

		key := duplicateKey(*row.Item)

		firstRow, exists := seen[key]

		if exists {
			findings = append(
				findings,
				validation.Finding{
					Row:      row.Row,
					Code:     "POSSIBLE_DUPLICATE",
					Severity: validation.SeverityWarning,
					Message: fmt.Sprintf(
						"possible duplicate of row %d",
						firstRow,
					),
				},
			)

			continue
		}

		seen[key] = row.Row
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

func AddFindings(
	rows []ProcessedRow,
	findings []validation.Finding,
) {

	for _, finding := range findings {

		for index := range rows {

			if rows[index].Row == finding.Row {

				rows[index].Findings = append(
					rows[index].Findings,
					finding,
				)

				break
			}
		}
	}
}
