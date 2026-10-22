package parser

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CSVResult struct {
	Headers []string
	Rows    [][]string
}

func ParseCSV(reader io.Reader) (*CSVResult, error) {
	csvReader := csv.NewReader(reader)

	headers, err := csvReader.Read()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to read CSV headers: %w",
			err,
		)
	}

	rows, err := csvReader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to read CSV rows: %w",
			err,
		)
	}

	return &CSVResult{
		Headers: headers,
		Rows:    rows,
	}, nil
}
