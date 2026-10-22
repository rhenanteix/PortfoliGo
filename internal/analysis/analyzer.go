package analysis

import (
	"time"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

type Analyzer struct {
	now time.Time
}

func NewAnalyzer(now time.Time) *Analyzer {
	return &Analyzer{
		now: now,
	}
}

func (a *Analyzer) Analyze(
	items []receivable.Receivable,
) PortfolioAnalysis {

	return PortfolioAnalysis{
		Summary: CalculateSummary(
			items,
			a.now,
		),
		Concentration: CalculateConcentration(
			items,
		),
		Aging: CalculateAging(
			items,
			a.now,
		),
	}
}
