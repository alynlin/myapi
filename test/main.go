/**
 * @Author: litsky
 * @Date:   2025/4/4
 * @Last Modified by:   litsky
 * @Last Modified time: 2025/4/4
 * @License: MIT
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var lock SpinLock
	wg := sync.WaitGroup{}

	var counter int
	// 启动多个协程模拟竞争
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			//wg.Add(1) 这里会有问题
			lock.Lock()
			defer wg.Done()
			counter++
			defer lock.UnLock()
		}()
	}

	wg.Wait()
	println("Counter:", counter) // 预期输出 Counter: 100

	ticket := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticket.C:
			fmt.Println(time.Now())
		}

	}

}

type SpinLock struct {
	State uint32
}

// 自旋锁
func (s *SpinLock) Lock() {

	for !atomic.CompareAndSwapUint32(&s.State, 0, 1) {
		runtime.Gosched()
	}

}

func (s *SpinLock) UnLock() {
	atomic.StoreUint32(&s.State, 0)
}
