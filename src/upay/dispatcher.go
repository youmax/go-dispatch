package main

import (
	"log"
	"upay/models"
)

type Dispatcher struct {
	WorkerPool models.WorkPool
	jobQueue   models.JobChannel
	maxWorkers uint
}

func NewDispatcher(maxWorkers uint) *Dispatcher {
	pool := make(models.WorkPool, maxWorkers)
	queue := make(models.JobChannel)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers, jobQueue: queue}
}

// 初始化worker池，并启动woker池，并开始接受新的job
func (d *Dispatcher) Run() {
	var i uint
	for ; i < d.maxWorkers; i++ {
		NewWorker(d.WorkerPool, i).Start()
	}
	go d.dispatch()
}

// 开始调度，接收新的job
func (d *Dispatcher) dispatch() {
	i := 0
	for {
		select {
		case job := <-d.jobQueue:
			go func(job models.Job) {
				log.Printf("Dispatcher get JOB")
				worker := <-d.WorkerPool
				worker <- job
				i++
				log.Printf("deal job %d", i)
			}(job)
		}
	}
}

func (d *Dispatcher) PushJob(job models.Job) {
	go func() {
		d.jobQueue <- job
	}()
}
