package models

type HealthResponse struct {
	Status string `json:"status"`

	QueueLength int `json:"queue_length"`
	QueueCap    int `json:"queue_capacity"`

	Submitted uint64 `json:"submitted_jobs"`
	Completed uint64 `json:"completed_jobs"`
	Failed    uint64 `json:"failed_jobs"`
}
