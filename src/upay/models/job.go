package models

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
