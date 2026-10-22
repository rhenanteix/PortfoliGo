package analysis

import (
	"testing"
	"time"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func TestCalculateSummary(t *testing.T) {

	now := time.Date(
		2026,
		9,
		20,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	items := []receivable.Receivable{
		{
			PayerDocument: "111",
			Amount:        100000,
			DueDate: now.AddDate(
				0,
				0,
				10,
			),
		},
		{
			PayerDocument: "222",
			Amount:        900000,
			DueDate: now.AddDate(
				0,
				0,
				50,
			),
		},
	}

	result := CalculateSummary(
		items,
		now,
	)

	if result.TotalAmount != 1000000 {
		t.Errorf(
			"expected total 1000000, got %d",
			result.TotalAmount,
		)
	}

	if result.Receivables != 2 {
		t.Errorf(
			"expected 2 receivables, got %d",
			result.Receivables,
		)
	}

	if result.UniquePayers != 2 {
		t.Errorf(
			"expected 2 unique payers, got %d",
			result.UniquePayers,
		)
	}

	if result.AverageTicket != 500000 {
		t.Errorf(
			"expected average ticket 500000, got %d",
			result.AverageTicket,
		)
	}

	if result.WeightedAverageTermDays != 46 {
		t.Errorf(
			"expected weighted term 46, got %d",
			result.WeightedAverageTermDays,
		)
	}
}
