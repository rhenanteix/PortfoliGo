package processing

import (
	"errors"

	"github.com/rhenanteix/goldscan/internal/normalization"
	"github.com/rhenanteix/goldscan/internal/receivable"
	"github.com/rhenanteix/goldscan/internal/validation"
)

type Processor struct {
	validationEngine *validation.Engine
}

func NewProcessor() *Processor {
	return &Processor{
		validationEngine: validation.NewEngine(),
	}
}

func (p *Processor) Process(
	inputs []receivable.Input,
) []ProcessedRow {

	results := make(
		[]ProcessedRow,
		0,
		len(inputs),
	)

	for index, input := range inputs {

		row := index + 2

		processed := ProcessedRow{
			Row: row,
			Raw: input,
		}

		item, err := normalization.NormalizeReceivable(input)

		if err != nil {
			processed.Findings = append(
				processed.Findings,
				normalizationErrorToFinding(
					row,
					err,
				),
			)

			results = append(results, processed)

			continue
		}

		processed.Item = &item

		itemFindings := p.validationEngine.ValidateItem(
			item,
			row,
		)

		processed.Findings = append(
			processed.Findings,
			itemFindings...,
		)

		results = append(results, processed)
	}

	return results
}

func normalizationErrorToFinding(
	row int,
	err error,
) validation.Finding {

	switch {

	case errors.Is(
		err,
		normalization.ErrInvalidAmount,
	):
		return validation.Finding{
			Row:      row,
			Code:     "INVALID_AMOUNT_FORMAT",
			Severity: validation.SeverityError,
			Field:    "amount",
			Message:  err.Error(),
		}

	case errors.Is(
		err,
		normalization.ErrInvalidDate,
	):
		return validation.Finding{
			Row:      row,
			Code:     "INVALID_DATE_FORMAT",
			Severity: validation.SeverityError,
			Message:  err.Error(),
		}

	default:
		return validation.Finding{
			Row:      row,
			Code:     "NORMALIZATION_ERROR",
			Severity: validation.SeverityError,
			Message:  err.Error(),
		}
	}
}

func ValidItems(
	rows []ProcessedRow,
) []receivable.Receivable {

	var items []receivable.Receivable

	for _, row := range rows {

		if row.Item == nil {
			continue
		}

		items = append(
			items,
			*row.Item,
		)
	}

	return items
}

func CollectFindings(
	rows []ProcessedRow,
) []validation.Finding {

	var findings []validation.Finding

	for _, row := range rows {

		findings = append(
			findings,
			row.Findings...,
		)
	}

	return findings
}
