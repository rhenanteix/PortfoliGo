package analysis

import (
	"time"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func CalculateAging(
	items []receivable.Receivable,
	now time.Time,
) Aging {

	var aging Aging

	today := startOfDay(now)

	for _, item := range items {

		days := int(
			item.DueDate.Sub(today).Hours() / 24,
		)

		switch {

		case days < 0:
			aging.Overdue += item.Amount

		case days <= 30:
			aging.Days0To30 += item.Amount

		case days <= 60:
			aging.Days31To60 += item.Amount

		case days <= 90:
			aging.Days61To90 += item.Amount

		default:
			aging.Over90Days += item.Amount
		}
	}

	return aging
}
