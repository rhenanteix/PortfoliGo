package analysis

import "github.com/rhenanteix/goldscan/internal/receivable"

func Total(items []receivable.Receivable) int64 {

	var total int64

	for _, item := range items {
		total += item.Amount
	}

	return total
}
