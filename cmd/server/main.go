package main

import (
	"log"
	"net/http"

	"github.com/Dharshan2208/code-compiler/internal/app"
	"github.com/Dharshan2208/code-compiler/internal/handler"
)

func main() {
	application := app.New()
	application.Pool.Start()

	http.HandleFunc("/run", handler.SubmitHandler(application))
	http.HandleFunc("/result/", handler.ResultHandler(application))
	http.HandleFunc("/health", handler.HealthHandler(application))

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
