package mapping

import "testing"

func TestValidateManualMapping(
	t *testing.T,
) {
	headers := []string{
		"RazaoSacado",
		"DocSacado",
		"NumeroTitulo",
		"ValorReceber",
		"DataPagamento",
	}

	columnMapping := ColumnMapping{
		"RazaoSacado":    FieldPayerName,
		"DocSacado":      FieldPayerDocument,
		"NumeroTitulo":   FieldInvoiceNumber,
		"ValorReceber":   FieldAmount,
		"DataPagamento":  FieldDueDate,
	}

	err := Validate(
		headers,
		columnMapping,
	)

	if err != nil {
		t.Fatalf(
			"expected valid mapping, got %v",
			err,
		)
	}
}

func TestValidateRejectsDuplicateFieldMapping(
	t *testing.T,
) {
	headers := []string{
		"ValorA",
		"ValorB",
		"Cliente",
		"Documento",
		"Vencimento",
	}

	columnMapping := ColumnMapping{
		"ValorA":     FieldAmount,
		"ValorB":     FieldAmount,
		"Cliente":    FieldPayerName,
		"Documento":  FieldPayerDocument,
		"Vencimento": FieldDueDate,
	}

	err := Validate(
		headers,
		columnMapping,
	)

	if err == nil {
		t.Fatal(
			"expected duplicate mapping error",
		)
	}
}