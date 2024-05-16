package logic

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

func print_in_turn2(s string, k int, param ...int) {
	ss := strings.Split(s, "")
	sourceCh := make(chan string, len(ss))
	go func() {
		defer close(sourceCh)
		for _, it := range ss {
			sourceCh <- it
		}
	}()
	sec2Use := 10
	if len(param) != 0 {
		sec2Use = param[0]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(sec2Use))
	defer cancel()
	wg := sync.WaitGroup{}
	startChs := make([]chan bool, 0, k)
	endChs := make([]chan bool, 0, k)
	for i := 0; i < k; i++ {
		startChs = append(startChs, make(chan bool))
		endChs = append(endChs, make(chan bool))
	}
	for i := 0; i < k; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					close(endChs[j])
					fmt.Printf("[%d]exit\n", j)
					return
				case <-startChs[j]:
					select {
					case c := <-sourceCh:
						fmt.Printf("[%d]%s\n", j, c)
					default:
					}
					select {
					case endChs[j] <- true:
					default:
					}
				}
			}
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// scheduler
		idx := 0
		for {
			select {
			case <-ctx.Done():
				for _, it := range startChs {
					close(it)
				}
				fmt.Printf("scheduler exit\n")
				return
			default:
				i := idx % k
				idx++
				startChs[i] <- true
				<-endChs[i]
			}
		}
	}()
	wg.Wait()
}
