package analysis

import (
	"time"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func CalculateSummary(
	items []receivable.Receivable,
	now time.Time,
) PortfolioSummary {

	var totalAmount int64

	payers := make(map[string]bool)

	for _, item := range items {

		totalAmount += item.Amount

		if item.PayerDocument != "" {
			payers[item.PayerDocument] = true
		}
	}

	var averageTicket int64

	if len(items) > 0 {
		averageTicket =
			totalAmount / int64(len(items))
	}

	return PortfolioSummary{
		TotalAmount:   totalAmount,
		Receivables:   len(items),
		UniquePayers:  len(payers),
		AverageTicket: averageTicket,
		WeightedAverageTermDays: weightedAverageTerm(
			items,
			now,
		),
	}
}

func weightedAverageTerm(
	items []receivable.Receivable,
	now time.Time,
) int {

	var weightedSum int64
	var totalAmount int64

	today := startOfDay(now)

	for _, item := range items {

		days := int64(
			item.DueDate.Sub(today).Hours() / 24,
		)

		if days < 0 {
			days = 0
		}

		weightedSum += item.Amount * days
		totalAmount += item.Amount
	}

	if totalAmount == 0 {
		return 0
	}

	return int(
		weightedSum / totalAmount,
	)
}

func startOfDay(value time.Time) time.Time {
	return time.Date(
		value.Year(),
		value.Month(),
		value.Day(),
		0,
		0,
		0,
		0,
		value.Location(),
	)
}
