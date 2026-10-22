package normalization

import "unicode"

func NormalizeDocument(value string) string {
	var result []rune

	for _, char := range value {
		if unicode.IsDigit(char) {
			result = append(result, char)
		}
	}

	return string(result)
}
