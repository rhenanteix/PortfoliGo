package normalization

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func NormalizeReceivable(
	input receivable.Input,
) (receivable.Receivable, error) {

	amount, err := ParseMoney(input.Amount)

	if err != nil {
		return receivable.Receivable{}, fmt.Errorf(
			"invalid amount: %w",
			err,
		)
	}

	dueDate, err := ParseDate(input.DueDate)

	if err != nil {
		return receivable.Receivable{}, fmt.Errorf(
			"invalid due date: %w",
			err,
		)
	}

	var issueDateValue time.Time

	if input.IssueDate != "" {
		issueDate, err := ParseDate(input.IssueDate)

		if err != nil {
			return receivable.Receivable{}, fmt.Errorf(
				"invalid issue date: %w",
				err,
			)
		}

		issueDateValue = issueDate
	}

	return receivable.Receivable{
		ID:            uuid.NewString(),
		PayerName:     NormalizeText(input.PayerName),
		PayerDocument: NormalizeDocument(input.PayerDocument),
		InvoiceNumber: NormalizeText(input.InvoiceNumber),
		Amount:        amount,
		IssueDate:     issueDateValue,
		DueDate:       dueDate,
	}, nil
}

func NormalizeAll(
	inputs []receivable.Input,
) ([]receivable.Receivable, error) {

	result := make(
		[]receivable.Receivable,
		0,
		len(inputs),
	)

	for index, input := range inputs {

		item, err := NormalizeReceivable(input)

		if err != nil {
			return nil, fmt.Errorf(
				"row %d: %w",
				index+2,
				err,
			)
		}

		result = append(result, item)
	}

	return result, nil
}
