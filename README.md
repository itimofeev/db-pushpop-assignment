# db-pushpop-assignment

```go
type DB interface { Push(data []byte) Pop() ([]byte, error)
}
type Storage interface {
    Push(ctx context.Context, data []byte)
	Pop(ctx context.Context) ([]byte, error)
}
```

Необходимо реализовать интерфейс Storage (DB интерфейс можно реализовать в качестве дополнительного задания). На вход он получает сообщения, отдавать их должен от новых к старым. Сохранение/получение данных из DB затратная операция, поэтому необходимо оптимизировать интенсивный ввод/вывод, при этом нельзя потерять данные больше, чем за 1 минуту. Считаем, что сохраненное в DB сообщение потеряться не может. Контексты могут быть отменены вызывающей стороной, вызывающая сторона должна понимать по какой причине была возвращена ошибка. Можно использовать только встроенные средства языка, никакие внешние БД использовать нельзя.

Вот заготовка кода, реализовать необходимо методы с паниками. Использовать Storage могут по-разному, приведен только один кейс, чтобы показать, что использование будет происходить в разных горутинах.

```go
package main
import ( "context"
    "fmt"
    "sync"
    "time"
)
type DB interface { Push(data []byte) Pop() ([]byte, error)
}
type dbMock struct { }
func newDBMock() *dbMock { return &dbMock{}
}
   
func (db *dbMock) Push(data []byte) { time.Sleep(time.Millisecond * 5) // Just to emulate DB delay
}
func (db *dbMock) Pop() ([]byte, error) { time.Sleep(time.Millisecond * 5) // Just to emulate DB delay return []byte("test"), nil
}

type Storage interface {
    Push(ctx context.Context, data []byte)
	Pop(ctx context.Context) ([]byte, error)
}
type storage struct { 
	db DB
}
func newStorage(db DB) *storage { 
	return &storage{db: db}
}
func (s *storage) Push(ctx context.Context, data []byte) { 
	panic("Implement me")
}
func (s *storage) Pop(ctx context.Context) ([]byte, error) { 
	panic("Implement me")
}

func main() {
    ctx, cancel := context.WithCancel(context.Background()) defer cancel()
    s := newStorage(newDBMock())
    wg := sync.WaitGroup{}
    const count = 1000
    wg.Add(count * 2)
    for i := 0; i < count; i++ {
    go func(i int) {
    s.Push(ctx, []byte(fmt.Sprintf("%d %v", i, time.Now()))) wg.Done()
    }(i) }
    for i := 0; i < count; i++ { go func() {
                _, _ = s.Pop(ctx)
                wg.Done()
            }()
    }
    wg.Wait() 
}
```