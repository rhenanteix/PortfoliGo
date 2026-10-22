package validation

import (
	"testing"
	"time"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func TestFindDuplicates(t *testing.T) {

	dueDate := time.Date(
		2026,
		10,
		15,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	items := []receivable.Receivable{
		{
			PayerDocument: "11444777000161",
			InvoiceNumber: "1001",
			Amount:        1250000,
			DueDate:       dueDate,
		},
		{
			PayerDocument: "11444777000161",
			InvoiceNumber: "1001",
			Amount:        1250000,
			DueDate:       dueDate,
		},
	}

	findings := FindDuplicates(items)

	if len(findings) != 1 {
		t.Fatalf(
			"expected 1 duplicate finding, got %d",
			len(findings),
		)
	}

	if findings[0].Row != 3 {
		t.Errorf(
			"expected duplicate on row 3, got row %d",
			findings[0].Row,
		)
	}
}
