package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/Dharshan2208/code-compiler/internal/app"
	"github.com/Dharshan2208/code-compiler/internal/models"
)

func SubmitHandler(application *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RunRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("submit request rejected: invalid json: %v", err)
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		jobID := uuid.New().String()

		job := &models.Job{
			ID:       jobID,
			Language: req.Language,
			Code:     req.Code,
			Status:   "pending",
		}

		application.Store.Add(job)
		application.Queue.Push(job)

		log.Printf("Job submitted: ID=%s Language=%s", job.ID, job.Language)

		response := models.SubmitResponse{
			JobID:  jobID,
			Status: "pending",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func ResultHandler(application *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/result/")
		log.Printf("Result requested: ID=%s", id)

		job, exists := application.Store.Get(id)
		if !exists {
			log.Printf("result request failed: id=%s reason=job_not_found", id)
			http.Error(w, "job not found", http.StatusNotFound)
			return
		}

		log.Printf("Result returned: ID=%s status=%s", job.ID, job.Status)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(job)
	}
}
