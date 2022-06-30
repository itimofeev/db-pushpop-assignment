package db

import "time"

type mock struct{}

func NewDBMock() *mock {
	return &mock{}
}

func (db *mock) Push(data []byte) {
	time.Sleep(time.Millisecond * 5) // Just to emulate DB delay
}
func (db *mock) Pop() ([]byte, error) {
	time.Sleep(time.Millisecond * 5) // Just to emulate DB delay
	return []byte("test"), nil
}
