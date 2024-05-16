package logic

import (
	"context"
	"fmt"
	"mytest/model"
	"strings"
	"sync"
	"time"
)

func print_in_turn3(s string, k int, param ...int) {
	ss := strings.Split(s, "")
	strIdx := 0
	// idLocker := new(sync.Mutex)
	sec2Use := 10
	if len(param) != 0 {
		sec2Use = param[0]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(sec2Use))
	defer cancel()
	tokenCh := make(chan model.Empty, 1)
	tokenCh <- model.Empty{}
	wg := sync.WaitGroup{}
	for i := 0; i < k; i++ {
		wg.Add(1)
		go func(gid int) {
			defer wg.Done()
			for {
				select {
				case t := <-tokenCh:
					charCurr := ss[strIdx%len(ss)]
					strIdx++
					if strIdx == len(ss) {
						strIdx = 0
					}
					fmt.Printf("[%d] %s\n", gid, charCurr)
					tokenCh <- t
				case <-ctx.Done():
					fmt.Printf("[%d] exit\n", gid)
					return
				}
			}
		}(i)
	}
	wg.Wait()
}
