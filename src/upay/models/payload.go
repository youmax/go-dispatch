package models

// Payload interface for polymorphism
type Payload interface {
	Handle()
}
