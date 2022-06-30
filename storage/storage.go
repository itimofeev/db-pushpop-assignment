package storage

import "context"

type DB interface {
	Push(data []byte)
	Pop() ([]byte, error)
}

type storage struct {
	db DB
}

func NewStorage(db DB) *storage {
	return &storage{db: db}
}
func (s *storage) Push(ctx context.Context, data []byte) {
	panic("Implement me")
}

func (s *storage) Pop(ctx context.Context) ([]byte, error) {
	panic("Implement me")
}
