package main

import (
	"log"
	"net/http"

	"github.com/jhonM8a/worker-evaluacion/api"
	"github.com/jhonM8a/worker-evaluacion/config"
	"github.com/jhonM8a/worker-evaluacion/internal/dispatcher"
	Job "github.com/jhonM8a/worker-evaluacion/internal/job"
)

func main() {
	// Cargar configuraciones
	cfg := config.Load()

	// Crear cola de trabajos
	jobQueue := make(chan Job.Job, cfg.MaxQueue)

	// Inicializar el dispatcher
	dispatcher := dispatcher.NewDispatcher(jobQueue, cfg.MaxWorkers)
	dispatcher.Run()

	// Configurar el servidor HTTP
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		api.RequestHandler(w, r, jobQueue)
	})

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
