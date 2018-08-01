package models

// Payload interface for polymorphism
type Payload interface {
	Handle()
}

// Job represents the job to be run
type Job struct {
	Url       string
	Payload   Payload
	Exception string
	Attempts  uint8
	DeletedAt string
	FailedAt  string
}

// JobChannel ...
type JobChannel chan Job

// WorkPool ...
type WorkPool chan JobChannel
