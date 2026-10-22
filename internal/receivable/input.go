package receivable

type Input struct {
	PayerName     string `json:"payer_name"`
	PayerDocument string `json:"payer_document"`
	InvoiceNumber string `json:"invoice_number"`
	Amount        string `json:"amount"`
	IssueDate     string `json:"issue_date"`
	DueDate       string `json:"due_date"`
}
