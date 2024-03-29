package lock

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type spinLock struct {
	owner int64
	count int
	flag  *uint32
}

func NewSpinLock() sync.Locker {
	var lock spinLock
	return &lock
}

func (sl *spinLock) Lock() {
	me := GetGoroutineId()
	if sl.owner == me { // 如果当前线程已经获取到了锁，线程数增加一，然后返回
		sl.count++
		return
	}
	// 如果没获取到锁，则通过CAS自旋
	for !atomic.CompareAndSwapUint32(sl.flag, 0, 1) {
		runtime.Gosched()
	}
}
func (sl *spinLock) Unlock() {
	if sl.owner != GetGoroutineId() {
		panic("illegalMonitorStateError")
	}
	if sl.count > 0 { // 如果大于0，表示当前线程多次获取了该锁，释放锁通过count减一来模拟
		sl.count--
	} else { // 如果count==0，可以将锁释放，这样就能保证获取锁的次数与释放锁的次数是一致的了。
		atomic.StoreUint32((*uint32)(sl.flag), 0)
	}
}

func GetGoroutineId() int64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return int64(id)
}
