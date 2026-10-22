package validation

func BuildSummary(
	total int,
	findings []Finding,
) Summary {

	rowsWithWarning := make(map[int]bool)
	rowsWithError := make(map[int]bool)

	for _, finding := range findings {

		switch finding.Severity {

		case SeverityError:
			rowsWithError[finding.Row] = true

		case SeverityWarning:
			if !rowsWithError[finding.Row] {
				rowsWithWarning[finding.Row] = true
			}
		}
	}

	for row := range rowsWithError {
		delete(rowsWithWarning, row)
	}

	invalid := len(rowsWithError)
	review := len(rowsWithWarning)

	valid := total - invalid - review

	return Summary{
		Total:   total,
		Valid:   valid,
		Review:  review,
		Invalid: invalid,
	}
}
