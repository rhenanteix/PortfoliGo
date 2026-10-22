package processing

import (
	"testing"

	"github.com/rhenanteix/goldscan/internal/receivable"
)

func TestProcessorContinuesAfterInvalidRow(
	t *testing.T,
) {

	inputs := []receivable.Input{
		{
			PayerName:     "Empresa A",
			PayerDocument: "11444777000161",
			InvoiceNumber: "1001",
			Amount:        "1000,00",
			DueDate:       "15/10/2026",
		},
		{
			PayerName:     "Empresa B",
			PayerDocument: "11444777000161",
			InvoiceNumber: "1002",
			Amount:        "ABC",
			DueDate:       "15/10/2026",
		},
		{
			PayerName:     "Empresa C",
			PayerDocument: "11444777000161",
			InvoiceNumber: "1003",
			Amount:        "3000,00",
			DueDate:       "15/10/2026",
		},
	}

	processor := NewProcessor()

	rows := processor.Process(inputs)

	if len(rows) != 3 {
		t.Fatalf(
			"expected 3 processed rows, got %d",
			len(rows),
		)
	}

	if rows[1].Item != nil {
		t.Error(
			"expected second row to fail normalization",
		)
	}

	if rows[2].Item == nil {
		t.Error(
			"expected third row to be processed",
		)
	}
}
