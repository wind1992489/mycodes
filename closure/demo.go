package main

import (
	"fmt"
	"sync"
)

type Counter func() int

func NewCounter() Counter {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// g := NewCounter()
	g := func() {
		fmt.Println("hello")
	}
	wg := sync.WaitGroup{}
	once := sync.Once{}
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func(idx int) {
			defer wg.Done()
			once.Do(g)
			// fmt.Printf("[%d]:%d\n", idx, g())
		}(i)
	}
	wg.Wait()
}
