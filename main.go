package main2

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//	func main1() {
//		s := []rune("hello world~")
//		wg := new(sync.WaitGroup)
//		wg.Add(2)
//		sCh := make(chan rune, len(s))
//		bCh := make(chan bool, 1)
//		for _, it := range s {
//			sCh <- it
//		}
//		// worker1
//		go func() {
//			defer wg.Done()
//			for {
//				f, ok := <-bCh
//				if !ok {
//					return
//				}
//				select {
//				case s1 := <-sCh:
//					fmt.Println("go1: ", string(s1))
//				default:
//					close(bCh)
//					return
//				}
//				bCh <- f
//			}
//		}()
//		// worker2
//		go func() {
//			defer wg.Done()
//			for {
//				f, ok := <-bCh
//				if !ok {
//					return
//				}
//				select {
//				case s1 := <-sCh:
//					fmt.Println("go2: ", string(s1))
//				default:
//					close(bCh)
//					return
//				}
//				bCh <- f
//			}
//		}()
//		bCh <- true
//		wg.Wait()
//	}
//
//	func main2() {
//		// 已经被关闭的channel可以再被读吗
//		ch := make(chan bool)
//		go func() {
//			ch <- true
//			time.Sleep(3 * time.Second)
//			close(ch)
//		}()
//		t := time.After(time.Second * 5)
//
// outer:
//
//		for {
//			select {
//			case x := <-ch:
//				fmt.Println("received:", x)
//			case <-t:
//				break outer
//			default:
//			}
//		}
//	}
//
//	func main5() {
//		msg := []string{"a", "b", "c"}
//		chs := make([]chan bool, 3)
//		hear := make(chan bool)
//		for i := 0; i < 3; i++ {
//			chs[i] = make(chan bool)
//			go func(ch chan bool, m string, h chan bool) {
//				j := 1
//				for {
//					select {
//					case f := <-ch:
//						if f {
//							fmt.Println(j, ":", m)
//							j++
//							hear <- true
//						} else {
//							return
//						}
//					}
//				}
//			}(chs[i], msg[i], hear)
//		}
//		for i := 0; i < 30; i++ {
//			chs[i%3] <- true
//			<-hear
//		}
//		for i := 0; i < len(chs); i++ {
//			chs[i] <- false
//		}
//	}
//
//	func main6() {
//		ch := make(chan int, 1) // Create a buffered channel to avoid blocking
//
//		// Close the channel
//		close(ch)
//
//		// Attempt to send data to the closed channel
//		select {
//		case ch <- 1:
//			fmt.Println("Sent data to channel")
//		default:
//			fmt.Println("Cannot send data to closed channel")
//		}
//	}
//
//	func main7() {
//		defer func() {
//			if r := recover(); r != nil {
//				fmt.Println(r)
//			}
//		}()
//		k := 3
//		chs := make([]chan model.Empty, k)
//		for i := 0; i < k; i++ {
//			chs[i] = make(chan model.Empty)
//		}
//		wt := new(sync.WaitGroup)
//		for i := 0; i < k; i++ {
//			wt.Add(1)
//			go func(idx int) {
//				defer func() {
//					fmt.Printf("g[%d] return\n", idx)
//					wt.Done()
//				}()
//				myCh := chs[idx]
//				nextCh := chs[(idx+1)%k]
//				for {
//					select {
//					case _, ok := <-myCh:
//						if ok {
//							fmt.Printf("g[%d]:hello\n", idx)
//							nextCh <- model.Empty{}
//						} else {
//							// closed
//							close(nextCh)
//							return
//						}
//					}
//				}
//			}(i)
//		}
//		chs[0] <- model.Empty{}
//		time.Sleep(5 * time.Second)
//		close(chs[0])
//		wt.Wait()
//	}
//
//	func main8() {
//		ctx, cancel := context.WithCancel(context.Background())
//		msgCh := make(chan string)
//		wg := new(sync.WaitGroup)
//		wg.Add(2)
//		go func() {
//			defer func() {
//				cancel()
//				close(msgCh)
//				wg.Done()
//			}()
//			for i := 0; i < 100; i++ {
//				msgCh <- strconv.Itoa(i)
//			}
//		}()
//		go func() {
//			defer wg.Done()
//			myPrint(ctx, msgCh, 0)
//		}()
//		wg.Wait()
//	}

func main23472() {
	bufferSize := 10
	producerNum := 5
	workerNum := 10
	// 生产者每隔50ms生成一个数字，worker每隔500ms从管道中取数字并打印
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	msgCh := make(chan string, bufferSize)
	wg := new(sync.WaitGroup)
	for i := 0; i < producerNum; i++ {
		wg.Add(1)
		go func(ctx context.Context, producerID int) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			t := time.NewTicker(time.Millisecond * 50)
			defer t.Stop()
			for {
				select {
				case <-t.C:
					select {
					case msgCh <- fmt.Sprintf("%d_%d", producerID, r.Intn(1000)):
					default:
					}
				case <-ctx.Done():
					fmt.Println("producer", producerID, "exit")
					return
				}
			}
		}(ctx, i)
	}
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go func(ctx context.Context, workerID int) {
			defer wg.Done()
			myPrint(ctx, msgCh, workerID)
		}(ctx, i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		cancel()
	}()
	wg.Wait()
}
func myPrint(ctx context.Context, msgCh chan string, id int) {
	t := time.NewTicker(time.Millisecond * 500)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			select {
			case s := <-msgCh:
				fmt.Println("rcvd: ", s)
			default:
			}
		case <-ctx.Done():
			fmt.Println("worker[", id, "] exit", time.Now())
			return
		}
	}
}

type tps int

const (
	x tps = iota
	y
	z
)

func main222() {
	fmt.Println(runtime.Caller(0))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	d1(ctx)
}
func d1(ctx context.Context) {
	fmt.Println(runtime.Caller(3))
	select {
	case <-ctx.Done():
		return
	default:
	}
	val := ctx.Value("cnt")
	if val == nil {
		val = 0
		ctx = context.WithValue(ctx, "cnt", val)
	}
	if val == 5 {
		return
	} else {
		ctx = context.WithValue(ctx, "cnt", val.(int)+1)
	}
	d1(ctx)
}
