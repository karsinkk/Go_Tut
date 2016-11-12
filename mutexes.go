package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}
	var ops int64 = 0
	var total int = 0

	for i := 0; i < 100; i++ {
		go func() {
			for {
				key := rand.Intn(25)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()

			}
		}()
	}

	for j := 0; j < 50; j++ {
		go func() {
			for {
				key := rand.Intn(25)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println(opsFinal, " : Operations have been performed!!")

	mutex.Lock()
	fmt.Println("State : ", state)
	fmt.Println("Total : ", total)
	mutex.Unlock()

}
