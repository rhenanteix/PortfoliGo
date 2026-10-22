package mapping

import (
	"fmt"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func MapRows(
	headers []string,
	rows [][]string,
	columnMapping ColumnMapping,
) ([]receivable.Input, error) {

	var result []receivable.Input

	for rowIndex, row := range rows {

		if len(row) != len(headers) {
			return nil, fmt.Errorf(
				"row %d has %d columns, expected %d",
				rowIndex+2,
				len(row),
				len(headers),
			)
		}

		var input receivable.Input

		for columnIndex, value := range row {

			header := headers[columnIndex]

			field, exists := columnMapping[header]

			if !exists {
				continue
			}

			setField(&input, field, value)
		}

		result = append(result, input)
	}

	return result, nil
}

func setField(
	input *receivable.Input,
	field Field,
	value string,
) {
	switch field {

	case FieldPayerName:
		input.PayerName = value

	case FieldPayerDocument:
		input.PayerDocument = value

	case FieldInvoiceNumber:
		input.InvoiceNumber = value

	case FieldAmount:
		input.Amount = value

	case FieldIssueDate:
		input.IssueDate = value

	case FieldDueDate:
		input.DueDate = value
	}
}
