package dispatcher

import (
	job "github.com/jhonM8a/worker-evaluacion/internal/job"
	"github.com/jhonM8a/worker-evaluacion/internal/worker"
)

type Dispatcher struct {
	WorkerPool chan chan job.Job
	MaxWorkers int
	JobQueue   chan job.Job
}

func NewDispatcher(jobQueue chan job.Job, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan job.Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: workerPool,
	}
}

func (d *Dispatcher) Distpatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		w := worker.NewWorker(i, d.WorkerPool)
		w.Start()
	}
	go d.Distpatch()
}
