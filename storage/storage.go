package storage

import (
	"context"
	"time"
)

type DB interface {
	Push(data []byte)
	Pop() ([]byte, error)
}

type storage struct {
	db DB
	// canLoseDuration duration in which we can lose messages
	canLoseDuration time.Duration
}

func NewStorage(db DB, canLoseDuration time.Duration) *storage {
	return &storage{
		db:              db,
		canLoseDuration: canLoseDuration,
	}
}
func (s *storage) Push(ctx context.Context, data []byte) {
	panic("Implement me")
}

func (s *storage) Pop(ctx context.Context) ([]byte, error) {
	panic("Implement me")
}
