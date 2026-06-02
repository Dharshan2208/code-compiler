package queue

import (
	"github.com/Dharshan2208/code-compiler/internal/models"
)

type Queue struct {
	Jobs chan *models.Job
}

func NewQueue(size int) *Queue {
	return &Queue{
		Jobs: make(chan *models.Job, size),
	}
}

func (q *Queue) Push(job *models.Job) {
	q.Jobs <- job
}

func (q *Queue) Pop() *models.Job {
	return <-q.Jobs
}
