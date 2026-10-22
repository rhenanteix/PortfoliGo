package mapping

type Field string

type ColumnMapping map[string]Field

const (
	FieldPayerName     Field = "payer_name"
	FieldPayerDocument Field = "payer_document"
	FieldInvoiceNumber Field = "invoice_number"
	FieldAmount        Field = "amount"
	FieldIssueDate     Field = "issue_date"
	FieldDueDate       Field = "due_date"
)

var RequiredFields = []Field{
	FieldPayerName,
	FieldPayerDocument,
	FieldAmount,
	FieldDueDate,
}

var AllFields = []Field{
	FieldPayerName,
	FieldPayerDocument,
	FieldInvoiceNumber,
	FieldAmount,
	FieldIssueDate,
	FieldDueDate,
}

func IsValidField(field Field) bool {
	for _, validField := range AllFields {

		if field == validField {
			return true
		}
	}

	return false
}