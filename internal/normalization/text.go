package normalization

import "strings"

func NormalizeText(value string) string {
	return strings.TrimSpace(value)
}
