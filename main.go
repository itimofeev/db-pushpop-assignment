package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/itimofeev/db-pushpop-assignment/db"
	"github.com/itimofeev/db-pushpop-assignment/storage"
)

type Storage interface {
	Push(ctx context.Context, data []byte)
	Pop(ctx context.Context) ([]byte, error)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := storage.NewStorage(db.NewDBMock())

	wg := sync.WaitGroup{}
	const count = 1000
	wg.Add(count * 2)

	for i := 0; i < count; i++ {
		go func(i int) {
			s.Push(ctx, []byte(fmt.Sprintf("%d %v", i, time.Now())))
			wg.Done()
		}(i)
	}
	for i := 0; i < count; i++ {
		go func() {
			_, _ = s.Pop(ctx)
			wg.Done()
		}()
	}
	wg.Wait()
}
