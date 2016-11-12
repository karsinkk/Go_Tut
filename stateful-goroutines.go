package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {

	var readOps int64 = 0
	var writeOps int64 = 0

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 50; i++ {
		go func() {
			for {
				write := &writeOp{
					key:   rand.Intn(25),
					value: rand.Intn(100),
					resp:  make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&writeOps, 1)
			}
		}()
	}

	for j := 0; j < 30; j++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(25),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&readOps, 1)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadInt64(&readOps)
	writeOpsFinal := atomic.LoadInt64(&writeOps)
	fmt.Println(writeOpsFinal, " : Writes Performed")
	fmt.Println(readOpsFinal, " : Reads Performed")
}
