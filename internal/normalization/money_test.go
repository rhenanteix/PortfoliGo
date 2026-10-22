package normalization

import "testing"

func TestParseMoney(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{
			name:     "Brazilian format",
			input:    "R$ 12.500,90",
			expected: 1250090,
		},
		{
			name:     "Decimal comma",
			input:    "12500,90",
			expected: 1250090,
		},
		{
			name:     "Decimal point",
			input:    "12500.90",
			expected: 1250090,
		},
		{
			name:     "Integer",
			input:    "12500",
			expected: 1250000,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			result, err := ParseMoney(test.input)

			if err != nil {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			if result != test.expected {
				t.Errorf(
					"expected %d, got %d",
					test.expected,
					result,
				)
			}
		})
	}
}
