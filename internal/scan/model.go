package scan

import "time"

type Status string

const (
	StatusCreated    Status = "created"
	StatusUploaded   Status = "uploaded"
	StatusProcessing Status = "processing"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
)

type Scan struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
