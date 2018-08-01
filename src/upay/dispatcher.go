package main

import (
	"fmt"

	"upay/models"
)

type Dispatcher struct {
	WorkerPool models.WorkPool
	maxWorkers int
}

var JobQueue models.JobChannel

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(models.WorkPool, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

// 初始化worker池，并启动woker池，并开始接受新的job
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

// 开始调度，接收新的job
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job models.Job) {
				fmt.Println("[UPYUN] Dispatcher get JOB")
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				// 从pool中获取空闲的job channel
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				// 将job塞入 job channel中
				jobChannel <- job
			}(job)
		}
	}
}
