package normalization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidAmount = errors.New("invalid amount")

func ParseMoney(value string) (int64, error) {
	value = strings.TrimSpace(value)

	value = strings.ReplaceAll(value, "R$", "")
	value = strings.ReplaceAll(value, " ", "")

	if value == "" {
		return 0, fmt.Errorf(
			"%w: invalid format %q",
			ErrInvalidAmount,
			value,
		)
	}

	normalized := normalizeDecimalSeparator(value)

	parts := strings.Split(normalized, ".")

	if len(parts) > 2 {
		return 0, fmt.Errorf(
			"%w: invalid format %q",
			ErrInvalidAmount,
			value,
		)
	}

	integerPart := parts[0]

	decimalPart := "00"

	if len(parts) == 2 {
		decimalPart = parts[1]
	}

	switch len(decimalPart) {
	case 0:
		decimalPart = "00"

	case 1:
		decimalPart += "0"

	case 2:

	default:
		return 0, fmt.Errorf(
			"%w: amount has more than 2 decimal places: %s",
			ErrInvalidAmount,
			value,
		)
	}

	integerValue, err := strconv.ParseInt(
		integerPart,
		10,
		64,
	)

	if err != nil {
		return 0, fmt.Errorf(
			"%w: invalid integer part %q: %w",
			ErrInvalidAmount,
			integerPart,
			err,
		)
	}

	decimalValue, err := strconv.ParseInt(
		decimalPart,
		10,
		64,
	)

	if err != nil {
		return 0, fmt.Errorf(
			"%w: invalid decimal part %q: %w",
			ErrInvalidAmount,
			decimalPart,
			err,
		)
	}

	return integerValue*100 + decimalValue, nil
}

func normalizeDecimalSeparator(value string) string {
	lastComma := strings.LastIndex(value, ",")
	lastDot := strings.LastIndex(value, ".")

	switch {
	case lastComma > lastDot:
		// Ex.: 12.500,90
		value = strings.ReplaceAll(value, ".", "")
		value = strings.Replace(value, ",", ".", 1)

	case lastDot > lastComma:
		// Ex.: 12,500.90
		value = strings.ReplaceAll(value, ",", "")

	case lastComma >= 0:
		// Ex.: 12500,90
		value = strings.Replace(value, ",", ".", 1)
	}

	return value
}
