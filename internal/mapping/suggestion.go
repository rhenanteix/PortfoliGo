package mapping

import "strings"

var aliases = map[Field][]string{
	FieldPayerName: {
		"sacado",
		"cliente",
		"razao social",
		"razão social",
	},

	FieldPayerDocument: {
		"cnpj",
		"cpf/cnpj",
		"documento",
		"cnpj sacado",
	},

	FieldInvoiceNumber: {
		"nf",
		"nota",
		"nota fiscal",
		"numero nf",
	},

	FieldAmount: {
		"valor",
		"vl titulo",
		"vltitulo",
		"valor titulo",
		"valor total",
	},

	FieldIssueDate: {
		"emissao",
		"emissão",
		"data emissao",
		"data emissão",
	},

	FieldDueDate: {
		"vencimento",
		"dt venc",
		"dtvenc",
		"data vencimento",
	},
}

func Suggest(headers []string) ColumnMapping {

	result := make(ColumnMapping)

	for _, header := range headers {

		normalizedHeader := normalizeHeader(header)

		for field, fieldAliases := range aliases {

			for _, alias := range fieldAliases {

				if normalizedHeader == normalizeHeader(alias) {
					result[header] = field
					break
				}
			}
		}
	}

	return result
}

func normalizeHeader(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)

	return value
}
