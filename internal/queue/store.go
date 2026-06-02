package queue

import (
	"sync"

	"github.com/Dharshan2208/code-compiler/internal/models"
)

type Store struct {
	Jobs map[string]*models.Job
	Mu   sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		Jobs: make(map[string]*models.Job),
	}
}

func (s *Store) Add(job *models.Job) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.Jobs[job.ID] = job
}

func (s *Store) Get(id string) (*models.Job, bool) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	job, exists := s.Jobs[id]

	return job, exists
}

func (s *Store) Update(job *models.Job) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.Jobs[job.ID] = job
}
