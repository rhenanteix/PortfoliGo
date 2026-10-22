package mapping

func Analyze(headers []string) Result {
	suggested := Suggest(headers)

	columns := make(
		[]ColumnSuggestion,
		0,
		len(headers),
	)

	unrecognized := make(
		[]string,
		0,
	)

	for _, header := range headers {
		field, exists := suggested[header]

		if !exists {
			columns = append(
				columns,
				ColumnSuggestion{
					SourceColumn: header,
					Recognized:   false,
				},
			)

			unrecognized = append(
				unrecognized,
				header,
			)

			continue
		}

		columns = append(
			columns,
			ColumnSuggestion{
				SourceColumn: header,
				Field:        field,
				Recognized:   true,
			},
		)
	}

	missingRequired :=
		findMissingRequiredFields(suggested)

	status := StatusReady

	if len(missingRequired) > 0 {
		status = StatusMappingRequired
	}

	return Result{
		Status:             status,
		Columns:            columns,
		Mapping:            suggested,
		Unrecognized:       unrecognized,
		MissingRequired:    missingRequired,
		RecognitionPercent: recognitionPercent(
			headers,
			suggested,
		),
	}
}
func findMissingRequiredFields(
	columnMapping ColumnMapping,
) []Field {

	mappedFields := make(
		map[Field]bool,
	)

	for _, field := range columnMapping {
		mappedFields[field] = true
	}

	var missing []Field

	for _, required := range RequiredFields {

		if !mappedFields[required] {
			missing = append(
				missing,
				required,
			)
		}
	}

	return missing
}

func recognitionPercent(
	headers []string,
	columnMapping ColumnMapping,
) float64 {

	if len(headers) == 0 {
		return 0
	}

	return float64(len(columnMapping)) /
		float64(len(headers)) *
		100
}