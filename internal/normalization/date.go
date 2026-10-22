package normalization

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrInvalidDate = errors.New("invalid date")

var supportedDateLayouts = []string{
	"2006-01-02",
	"02/01/2006",
	"02-01-2006",
}

func ParseDate(value string) (time.Time, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return time.Time{}, fmt.Errorf("date is empty")
	}

	for _, layout := range supportedDateLayouts {
		parsed, err := time.Parse(layout, value)

		if err == nil {
			return parsed, nil
		}
	}

	return time.Time{}, fmt.Errorf(
		"unsupported date format: %s",
		value,
	)
}
