package main

import (
	"os"

	"upay/models"
)

var (
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue  = os.Getenv("MAX_QUEUE")
)

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool models.WorkPool
	JobChannel models.JobChannel
	quit       chan bool
}

// NewWorker ...
func NewWorker(workerPool models.WorkPool) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(models.JobChannel), // 分配器中，会将任务交给jobChannel，下面会从这里读取到job
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel: // 当闲置的jobChannel中有job时
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
