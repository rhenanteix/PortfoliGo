package validation

import "github.com/rhenanteix/goldscan/internal/receivable"

type DocumentValidator struct{}

func (DocumentValidator) Validate(
	item receivable.Receivable,
	row int,
) []Finding {

	document := item.PayerDocument

	if document == "" {
		return nil
	}

	valid := false

	switch len(document) {
	case 11:
		valid = isValidCPF(document)

	case 14:
		valid = isValidCNPJ(document)

	default:
		valid = false
	}

	if valid {
		return nil
	}

	return []Finding{
		{
			Row:      row,
			Code:     "INVALID_DOCUMENT",
			Severity: SeverityError,
			Field:    "payer_document",
			Message:  "payer document is invalid",
		},
	}
}

func isValidCNPJ(document string) bool {
	if len(document) != 14 {
		return false
	}

	if allDigitsEqual(document) {
		return false
	}

	firstDigit := calculateCNPJDigit(
		document[:12],
		[]int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
	)

	if int(document[12]-'0') != firstDigit {
		return false
	}

	secondDigit := calculateCNPJDigit(
		document[:13],
		[]int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2},
	)

	return int(document[13]-'0') == secondDigit
}

func calculateCNPJDigit(
	base string,
	weights []int,
) int {

	sum := 0

	for index, char := range base {
		digit := int(char - '0')

		sum += digit * weights[index]
	}

	remainder := sum % 11

	if remainder < 2 {
		return 0
	}

	return 11 - remainder
}

func isValidCPF(document string) bool {
	if len(document) != 11 {
		return false
	}

	if allDigitsEqual(document) {
		return false
	}

	firstDigit := calculateCPFDigit(
		document[:9],
		10,
	)

	if int(document[9]-'0') != firstDigit {
		return false
	}

	secondDigit := calculateCPFDigit(
		document[:10],
		11,
	)

	return int(document[10]-'0') == secondDigit
}

func calculateCPFDigit(
	base string,
	initialWeight int,
) int {

	sum := 0
	weight := initialWeight

	for _, char := range base {
		digit := int(char - '0')

		sum += digit * weight

		weight--
	}

	remainder := sum % 11

	if remainder < 2 {
		return 0
	}

	return 11 - remainder
}

func allDigitsEqual(document string) bool {
	if len(document) == 0 {
		return false
	}

	firstDigit := document[0]

	for i := 1; i < len(document); i++ {
		if document[i] != firstDigit {
			return false
		}
	}

	return true
}
