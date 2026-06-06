package models

import "time"

type Job struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Code     string `json:"code"`
	Status   string `json:"status"`

	CreatedAt   time.Time `json:"created_at"`
	ClaimedAt   time.Time `json:"claimed_at"`
	CompletedAt time.Time `json:"completed_at"`

	Result RunResponse `json:"result"`
}

type JobResponse struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Status   string `json:"status"`

	CreatedAt   time.Time `json:"created_at"`
	ClaimedAt   time.Time `json:"claimed_at"`
	CompletedAt time.Time `json:"completed_at"`

	Result RunResponse `json:"result"`
}

func NewJobResponse(job *Job) JobResponse {
	return JobResponse{
		ID:          job.ID,
		Language:    job.Language,
		Status:      job.Status,
		CreatedAt:   job.CreatedAt,
		ClaimedAt:   job.ClaimedAt,
		CompletedAt: job.CompletedAt,
		Result:      job.Result,
	}
}
