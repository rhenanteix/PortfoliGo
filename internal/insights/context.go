package insights

import (
	"github.com/rhenanteix/goldscan/internal/analysis"
	"github.com/rhenanteix/goldscan/internal/validation"
)

type Context struct {
	Portfolio   analysis.PortfolioAnalysis
	DataQuality validation.Summary
}
