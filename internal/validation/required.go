package validation

import (
	"strings"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

type RequiredValidator struct{}

func (RequiredValidator) Validate(
	item receivable.Receivable,
	row int,
) []Finding {

	var findings []Finding

	if strings.TrimSpace(item.PayerName) == "" {
		findings = append(findings, Finding{
			Row:      row,
			Code:     "MISSING_PAYER_NAME",
			Severity: SeverityError,
			Field:    "payer_name",
			Message:  "payer name is required",
		})
	}

	if strings.TrimSpace(item.PayerDocument) == "" {
		findings = append(findings, Finding{
			Row:      row,
			Code:     "MISSING_PAYER_DOCUMENT",
			Severity: SeverityError,
			Field:    "payer_document",
			Message:  "payer document is required",
		})
	}

	if strings.TrimSpace(item.InvoiceNumber) == "" {
		findings = append(findings, Finding{
			Row:      row,
			Code:     "MISSING_INVOICE_NUMBER",
			Severity: SeverityWarning,
			Field:    "invoice_number",
			Message:  "invoice number is missing",
		})
	}

	return findings
}
