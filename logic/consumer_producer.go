package logic

import (
	"context"
	"fmt"
	"mytest/model"
	"sync"
	"time"
)

// producer -> consumer

type PcModel struct {
	ctx         context.Context
	producerNum int
	consumerNum int
	flowSize    int
	timeOut     int
	stopCh      chan model.Empty
}

func NewPcModel(ctx context.Context, producerNum int, consumerNum int, flowSize int, timeOut int) *PcModel {
	m := &PcModel{
		ctx:         ctx,
		producerNum: producerNum,
		consumerNum: consumerNum,
		flowSize:    flowSize,
		timeOut:     timeOut,
		stopCh:      make(chan model.Empty),
	}

	return m
}

// todo sync.Pool
func (m *PcModel) Run() {
	flowCh := make(chan string, m.flowSize)
	delayCh := make(chan string, m.flowSize*100)
	ctx, cancel := context.WithTimeout(m.ctx, time.Second*time.Duration(m.timeOut))
	defer cancel()
	defer func() {
		close(flowCh)
		close(delayCh)
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < m.producerNum; i++ {
		wg.Add(1)
		go func(ID int) {
			defer wg.Done()
			t := time.NewTicker(time.Second)
			defer t.Stop()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("producer[%d] exit\n", ID)
					return
				case <-t.C:
					msg := fmt.Sprintf("producer[%d] talk [%d]\n", ID, time.Now().Unix())
					select {
					case flowCh <- msg:
					default:
						fmt.Printf("⚠️⚠️⚠️{%s} might be lost\n", msg)
						msg = fmt.Sprintf("[delay] %s", msg)
						// todo 管道拥塞，放入延迟队列
						delayCh <- msg
					}
				}
			}
		}(i + 1)
	}
	for i := 0; i < m.consumerNum; i++ {
		wg.Add(1)
		go func(ID int) {
			defer wg.Done()
			// for msg := range flowCh {
			// 	fmt.Printf("consumer[%d] recvd %s\n", ID, msg)
			// }
			for {
				select {
				case msg := <-flowCh:
					fmt.Printf("consumer[%d] recvd %s\n", ID, msg)
				case msg := <-delayCh:
					fmt.Println(msg)
				case <-time.After(time.Second * 5):
					// 5s读不到则退出
					fmt.Printf("consumer[%d] exit\n", ID)
					return
				}
			}
		}(i)
	}
	wg.Wait()
}

func (m *PcModel) Stop() {
	// todo
}
