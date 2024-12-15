package domain

import "github.com/google/uuid"

// AggregateRoot represents a base entity with an ID.
type AggregateRoot struct {
	Id uuid.UUID
}

// AggregateRootMethod defines methods for handling aggregate roots.
type AggregateRootMethod interface {
	SetEvent(aggregateID uuid.UUID)
	GetID() uuid.UUID
}

func (a *AggregateRoot) GetID() uuid.UUID {
	return a.Id
}
