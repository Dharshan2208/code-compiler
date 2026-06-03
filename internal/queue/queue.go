package queue

import (
	"log"

	"github.com/Dharshan2208/code-compiler/internal/models"
)

type Queue struct {
	Jobs chan *models.Job
}

func NewQueue(size int) *Queue {
	log.Printf("creating job queue: size=%d", size)

	return &Queue{
		Jobs: make(chan *models.Job, size),
	}
}

func (q *Queue) Push(job *models.Job) {
	log.Printf("queue push: job_id=%s status=%s language=%s", job.ID, job.Status, job.Language)
	q.Jobs <- job
}

func (q *Queue) TryPush(job *models.Job) bool {
	select {
	case q.Jobs <- job:
		log.Printf("queue push: job_id=%s status=%s language=%s", job.ID, job.Status, job.Language)
		return true
	default:
		log.Printf("queue full: rejected job_id=%s language=%s length=%d capacity=%d", job.ID, job.Language, len(q.Jobs), cap(q.Jobs))
		return false
	}
}

func (q *Queue) Pop() *models.Job {
	job := <-q.Jobs
	log.Printf("queue pop: job_id=%s status=%s language=%s", job.ID, job.Status, job.Language)

	return job
}
