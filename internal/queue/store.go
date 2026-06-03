package queue

import (
	"log"
	"sync"

	"github.com/Dharshan2208/code-compiler/internal/models"
)

type Store struct {
	Jobs map[string]*models.Job
	Mu   sync.RWMutex
}

func NewStore() *Store {
	log.Println("creating in-memory job store")

	return &Store{
		Jobs: make(map[string]*models.Job),
	}
}

func (s *Store) Add(job *models.Job) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.Jobs[job.ID] = job
	log.Printf("store add: job_id=%s status=%s language=%s", job.ID, job.Status, job.Language)
}

func (s *Store) Get(id string) (*models.Job, bool) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	job, exists := s.Jobs[id]
	if exists {
		log.Printf("store get: job_id=%s status=%s found=true", id, job.Status)
	} else {
		log.Printf("store get: job_id=%s found=false", id)
	}

	return job, exists
}

func (s *Store) Update(job *models.Job) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.Jobs[job.ID] = job
	log.Printf("store update: job_id=%s status=%s language=%s", job.ID, job.Status, job.Language)
}
