package scan

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rhenanteix/goldscan/internal/analysis"
	"github.com/rhenanteix/goldscan/internal/insights"
	"github.com/rhenanteix/goldscan/internal/mapping"
	"github.com/rhenanteix/goldscan/internal/parser"
	"github.com/rhenanteix/goldscan/internal/processing"
	"github.com/rhenanteix/goldscan/internal/validation"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	item := Scan{
		ID:        uuid.NewString(),
		Name:      "Nova análise",
		Status:    StatusCreated,
		CreatedAt: time.Now(),
	}
	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(item)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		http.Error(
			w,
			"invalid multipart form",
			http.StatusBadRequest,
		)
		return
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(
			w,
			"file is required",
			http.StatusBadRequest,
		)
		return
	}

	defer file.Close()

	csvResult, err := parser.ParseCSV(file)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	mappingResult := mapping.Analyze(
	csvResult.Headers,
)

if mappingResult.Status ==
	mapping.StatusMappingRequired {

	response := struct {
		Filename string         `json:"filename"`
		Rows     int            `json:"rows"`
		Mapping  mapping.Result `json:"mapping"`
	}{
		Filename: header.Filename,
		Rows:     len(csvResult.Rows),
		Mapping:  mappingResult,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).
		Encode(response); err != nil {

		log.Printf(
			"failed to encode mapping response: %v",
			err,
		)
	}

	return
}

inputs, err := mapping.MapRows(
	csvResult.Headers,
	csvResult.Rows,
	mappingResult.Mapping,
)

if err != nil {
	http.Error(
		w,
		err.Error(),
		http.StatusBadRequest,
	)

	return
}


	processor := processing.NewProcessor()

	processedRows := processor.Process(inputs)

	analyzer := analysis.NewAnalyzer(
		time.Now(),
	)

	portfolioAnalysis := analyzer.Analyze(
		processing.ValidItems(processedRows),
	)

	duplicateFindings := processing.FindDuplicates(
		processedRows,
	)

	processing.AddFindings(
		processedRows,
		duplicateFindings,
	)

	findings := processing.CollectFindings(
		processedRows,
	)

	summary := validation.BuildSummary(
		len(processedRows),
		findings,
	)

	insightsEngine := insights.NewEngine(
		insights.DefaultConfig(),
	)

	generatedInsights := insightsEngine.Evaluate(
		insights.Context{
			Portfolio:   portfolioAnalysis,
			DataQuality: summary,
		},
	)
	response := struct {
		Filename  string                     `json:"filename"`
		Headers   []string                   `json:"headers"`
		Rows      int                        `json:"rows"`
		Columns   int                        `json:"columns"`
		Mapping   mapping.ColumnMapping      `json:"mapping"`
		Summary   validation.Summary         `json:"data_quality"`
		Findings  []validation.Finding       `json:"findings"`
		Portfolio analysis.PortfolioAnalysis `json:"portfolio"`
		Data      []processing.ProcessedRow  `json:"data"`
		Insights  []insights.Insight         `json:"insights"`
	}{
		Filename:  header.Filename,
		Headers:   csvResult.Headers,
		Rows:      len(csvResult.Rows),
		Columns:   len(csvResult.Headers),
		Mapping:   mappingResult.Mapping,
		Summary:   summary,
		Findings:  findings,
		Portfolio: portfolioAnalysis,
		Data:      processedRows,
		Insights:  generatedInsights,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
}
