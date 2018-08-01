package main

import (
	"log"
	"os"

	"upay/models"
)

var (
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue  = os.Getenv("MAX_QUEUE")
)

// Worker represents the worker that executes the job
type Worker struct {
	ID         uint
	WorkerPool models.WorkPool
	JobChannel models.JobChannel
	quit       chan bool
}

// NewWorker ...
func NewWorker(workerPool models.WorkPool, id uint) Worker {
	return Worker{
		ID:         id,
		WorkerPool: workerPool,
		JobChannel: make(models.JobChannel), // 分配器中，会将任务交给jobChannel，下面会从这里读取到job
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		log.Printf("worker %d is startup ", w.ID)
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel: // 当闲置的jobChannel中有job时
				// log.Printf("worker %d get job %s", w.ID, job.Payload)
				job.Payload.Handle()
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
