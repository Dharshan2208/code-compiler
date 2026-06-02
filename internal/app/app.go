package app

import "github.com/Dharshan2208/code-compiler/internal/queue"

type App struct {
	Queue *queue.Queue
	Store *queue.Store
}

func New() *App {
	return &App{
		Queue: queue.NewQueue(100),
		Store: queue.NewStore(),
	}
}
