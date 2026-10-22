package normalization

import "testing"

func TestParseDate(t *testing.T) {

	tests := []struct {
		input string
	}{
		{input: "2026-10-15"},
		{input: "15/10/2026"},
		{input: "15-10-2026"},
	}

	for _, test := range tests {

		_, err := ParseDate(test.input)

		if err != nil {
			t.Errorf(
				"expected %s to be valid: %v",
				test.input,
				err,
			)
		}
	}
}
