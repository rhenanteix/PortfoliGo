package mapping

import "testing"

func TestAnalyzeRequiresMappingWhenRequiredFieldsMissing(
	t *testing.T,
) {
	headers := []string{
		"RazaoSacado",
		"DocSacado",
		"NumeroTitulo",
		"ValorReceber",
		"DataPagamento",
	}

	result := Analyze(headers)

	if result.Status != StatusMappingRequired {
		t.Errorf(
			"expected status %s, got %s",
			StatusMappingRequired,
			result.Status,
		)
	}

	if len(result.MissingRequired) == 0 {
		t.Error(
			"expected missing required fields",
		)
	}
}

func TestAnalyzeReadyWhenRequiredFieldsAreMapped(
	t *testing.T,
) {
	headers := []string{
		"Cliente",
		"Documento",
		"Nota",
		"Valor",
		"Vencimento",
	}

	result := Analyze(headers)

	if result.Status != StatusReady {
		t.Errorf(
			"expected status %s, got %s",
			StatusReady,
			result.Status,
		)
	}

	if len(result.MissingRequired) != 0 {
		t.Errorf(
			"expected no missing required fields, got %d",
			len(result.MissingRequired),
		)
	}
}