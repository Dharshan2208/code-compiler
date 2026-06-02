package models

type Job struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Code     string `json:"-"`
	Status   string `json:"status"`

	Result RunResponse `json:"result"`
}
