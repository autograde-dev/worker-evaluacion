package worker

import (
	"fmt"

	job "github.com/jhonM8a/worker-evaluacion/internal/job"
	evaluation "github.com/jhonM8a/worker-evaluacion/pkg"
)

type Worker struct {
	Id         int
	JobQueue   chan job.Job
	WorkerPool chan chan job.Job
	QuitChan   chan bool
}

func NewWorker(id int, workerPool chan chan job.Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan job.Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker con id %d Iniciado\n", w.Id)
				evaluation.Evaluate(job.IDEValuation, job.NameFileAnswer, job.NameFileEvaluation, job.NameBucket)
			case <-w.QuitChan:
				fmt.Printf("Worker %d finalizado\n", w.Id)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
