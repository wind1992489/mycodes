package concurrent

import (
	"context"
	"fmt"
	"mytest/model"
	"sync"
	"testing"
	"time"
)

// 3个goroutine轮流打印
func TestPrintByTurn(t *testing.T) {
	k := 3
	wg := sync.WaitGroup{}
	token := make(chan model.Empty, 1)
	token <- model.Empty{}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	defer close(token)
	for i := 0; i < k; i++ {
		wg.Add(1)
		go func(ctx context.Context, id int) {
			defer wg.Done()
			for {
				select {
				case tmp := <-token:
					defer func() {
						token <- tmp
					}()
					fmt.Println("hello: ", id)
				case <-ctx.Done():
					fmt.Println("bye: ", id)
					return
				default:
					time.Sleep(time.Millisecond * 50)
				}
			}
		}(ctx, i)
	}
	wg.Wait()
}
