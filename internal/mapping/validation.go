package mapping

import "fmt"

func Validate(
	
	headers []string,
	columnMapping ColumnMapping,
) error {

	headerSet := make(
		map[string]bool,
		len(headers),
	)

	for _, header := range headers {
		headerSet[header] = true
	}

	for sourceColumn := range columnMapping {

		if !headerSet[sourceColumn] {
			return fmt.Errorf(
				"column %q does not exist in file",
				sourceColumn,
			)
		}
	}

	missing :=
		findMissingRequiredFields(
			columnMapping,
		)

	if len(missing) > 0 {
		return fmt.Errorf(
			"required fields are missing: %v",
			missing,
		)
	}

	usedFields := make(
		map[Field]string,
	)

	for sourceColumn, field :=
		range columnMapping {

		if previousColumn, exists :=
			usedFields[field]; exists {

			return fmt.Errorf(
				"field %q is mapped by both %q and %q",
				field,
				previousColumn,
				sourceColumn,
			)
		}

		if !IsValidField(field) {
		return fmt.Errorf(
			"invalid field %q",
			field,
		)
	}

	usedFields[field] = sourceColumn
}

	return nil

	
	
}