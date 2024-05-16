package lock

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGetGoroutineId(t *testing.T) {
	go func() {
		fmt.Println("sub:", GetGoroutineId())
	}()
	fmt.Println("main:", GetGoroutineId())
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
}
