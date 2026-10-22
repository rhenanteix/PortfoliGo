package validation

import (
	"time"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

type Validator interface {
	Validate(
		item receivable.Receivable,
		row int,
	) []Finding
}

type Engine struct {
	validators []Validator
}

func NewEngine() *Engine {
	return &Engine{
		validators: []Validator{
			RequiredValidator{},
			AmountValidator{},
			DocumentValidator{},
			DueDateValidator{
				Now: time.Now(),
			},
		},
	}
}

func (e *Engine) Validate(
	items []receivable.Receivable,
) []Finding {

	var findings []Finding

	for index, item := range items {

		row := index + 2

		findings = append(
			findings,
			e.ValidateItem(item, row)...,
		)
	}

	findings = append(
		findings,
		FindDuplicates(items)...,
	)

	return findings
}

func (e *Engine) ValidateItem(
	item receivable.Receivable,
	row int,
) []Finding {

	var findings []Finding

	for _, validator := range e.validators {

		itemFindings := validator.Validate(
			item,
			row,
		)

		findings = append(
			findings,
			itemFindings...,
		)
	}

	return findings
}
