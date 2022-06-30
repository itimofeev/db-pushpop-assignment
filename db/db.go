package db

import "time"

type mock struct {
	delay time.Duration
}

func NewDBMock(delay time.Duration) *mock {
	return &mock{
		delay: delay,
	}
}

func (db *mock) Push(data []byte) {
	time.Sleep(db.delay) // Just to emulate DB delay
}
func (db *mock) Pop() ([]byte, error) {
	time.Sleep(db.delay) // Just to emulate DB delay
	return []byte("test"), nil
}
