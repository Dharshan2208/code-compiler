package app

import (
	"log"

	"github.com/Dharshan2208/code-compiler/internal/queue"
	"github.com/Dharshan2208/code-compiler/internal/worker"
)

type App struct {
	Queue *queue.Queue
	Store *queue.Store
	Pool  *worker.Pool

	Stats *queue.Stats
}

func New() *App {
	log.Println("Initializing application...")

	q := queue.NewQueue(100)
	s := queue.NewStore()
	stats := &queue.Stats{}
	
	p := worker.NewPool(4, q, s, stats)

	log.Println("Application initialized with queue_size=100 worker_count=4")

	return &App{
		Queue: q,
		Store: s,
		Pool:  p,
		Stats: stats,
	}
}
