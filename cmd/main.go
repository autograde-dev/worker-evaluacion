package main

import (
	"log"
	"net/http"

	// "github.com/jhonM8a/worker-evaluacion/api"
	"github.com/jhonM8a/worker-evaluacion/config"
	"github.com/jhonM8a/worker-evaluacion/internal/dispatcher"
	Job "github.com/jhonM8a/worker-evaluacion/internal/job"
	rabittmq "github.com/jhonM8a/worker-evaluacion/internal/rabittmq"
)

func main() {

	// Cargar configuraciones
	cfg := config.Load()

	// Crear cola de trabajos
	jobQueue := make(chan Job.Job, cfg.MaxQueue)

	// Inicializar el dispatcher
	dispatcher := dispatcher.NewDispatcher(jobQueue, cfg.MaxWorkers)
	dispatcher.Run()

	// Escuchar mensajes de la cola "evaluation"
	err := rabittmq.ConsumeMessages(jobQueue)
	if err != nil {
		log.Fatalf("Error al consumir mensajes: %v", err)
	}

	// Configurar el servidor HTTP
	//http.HandleFunc("/eva", func(w http.ResponseWriter, r *http.Request) {
	//		api.RequestHandler(w, r, jobQueue, nil)
	//	})

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
