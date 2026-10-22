package analysis

import (
	"sort"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

type payerAggregate struct {
	Document string
	Name     string
	Amount   int64
}

func CalculateConcentration(
	items []receivable.Receivable,
) []PayerConcentration {

	payers := make(
		map[string]*payerAggregate,
	)

	var totalAmount int64

	for _, item := range items {

		totalAmount += item.Amount

		payer, exists :=
			payers[item.PayerDocument]

		if !exists {

			payer = &payerAggregate{
				Document: item.PayerDocument,
				Name:     item.PayerName,
			}

			payers[item.PayerDocument] = payer
		}

		payer.Amount += item.Amount
	}

	result := make(
		[]PayerConcentration,
		0,
		len(payers),
	)

	for _, payer := range payers {

		var percentage float64

		if totalAmount > 0 {
			percentage =
				float64(payer.Amount) /
					float64(totalAmount) *
					100
		}

		result = append(
			result,
			PayerConcentration{
				PayerDocument: payer.Document,
				PayerName:     payer.Name,
				Amount:        payer.Amount,
				Percentage:    percentage,
			},
		)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Amount >
				result[j].Amount
		},
	)

	return result
}
