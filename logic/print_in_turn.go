package logic

import (
	"fmt"
	"log"
	"strings"
)

// 将输入字符串按照指定的协程数目轮流打印
func PrintInTurn(input string, k ...int) error {
	if len(input) == 0 {
		return nil
	}
	charList := strings.Split(input, "")
	charCh := make(chan string, len(charList))
	for _, it := range charList {
		charCh <- it
	}
	// 参与打印协程数
	var n int
	if len(k) == 0 {
		n = 2
	} else {
		n = k[0]
	}
	workerList := make([]*worker, n)
	for i := 0; i < n; i++ {
		stuff := newWorker(i + 1)
		stuff.task = func() {
			select {
			case char, ok := <-charCh:
				if !ok {
					// charCh is closed, release all workers
					stuff.stopCh <- true
					return
				}
				fmt.Printf("go{%d}:%s\n", stuff.id, char)
			}
		}
		go stuff.doJob()
		workerList[i] = stuff
	}
	return nil
}

type worker struct {
	id     int
	workCh chan bool
	stopCh chan bool
	task   func()
}

func newWorker(id int) *worker {
	return &worker{
		id:     id,
		workCh: make(chan bool),
		stopCh: make(chan bool),
		task:   nil,
	}
}
func (w *worker) doJob() {
	for {
		select {
		case <-w.workCh:
			w.task()
		case <-w.stopCh:
			log.Println("worker{", w.id, "}finished job")
			return
		}
	}
}
