package app

import (
	"log"

	"github.com/Dharshan2208/code-compiler/internal/queue"
	redisclient "github.com/Dharshan2208/code-compiler/internal/redis"
	"github.com/Dharshan2208/code-compiler/internal/store"
	"github.com/Dharshan2208/code-compiler/internal/worker"
)

type App struct {
	Queue *queue.Queue
	Store *store.RedisStore
	Pool  *worker.Pool

	Stats *queue.Stats
}

func New() *App {
	return NewWorker()
}

func NewAPI() *App {
	return newApp("api", 0)
}

func NewWorker() *App {
	return newApp("worker", 4)
}

func newApp(role string, workerCount int) *App {
	log.Printf("application initializing: role=%s", role)

	redisClient := redisclient.New()
	q := queue.NewQueue(redisClient, 100)
	s := store.NewRedisStore(redisClient)
	stats := &queue.Stats{}

	var p *worker.Pool
	if workerCount > 0 {
		p = worker.NewPool(workerCount, q, s, stats)
	}

	log.Printf("application initialized: role=%s queue_size=100 worker_count=%d", role, workerCount)

	return &App{
		Queue: q,
		Store: s,
		Pool:  p,
		Stats: stats,
	}
}
