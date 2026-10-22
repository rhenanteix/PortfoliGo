package receivable

import "time"

type Receivable struct {
	ID            string    `json:"id"`
	PayerName     string    `json:"payer_name"`
	PayerDocument string    `json:"payer_document"`
	InvoiceNumber string    `json:"invoice_number"`
	Amount        int64     `json:"amount"`
	IssueDate     time.Time `json:"issue_date"`
	DueDate       time.Time `json:"due_date"`
}
