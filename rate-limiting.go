package main

import (
	"fmt"
	"time"
)

func main() {

	req := make(chan int, 5)
	for i := 1; i < 6; i++ {
		req <- i
	}
	close(req)

	limit := time.Tick(time.Second * 2)

	for r := range req {
		<-limit
		fmt.Println("Request ", r, " processed ", "at :", time.Now())
	}

	fmt.Println()

	bufferLimiter := make(chan time.Time, 3)

	for j := 0; j < 3; j++ {
		bufferLimiter <- time.Now()
	}

	go func() {
		for s := range time.NewTicker(time.Millisecond * 200).C {
			bufferLimiter <- s
		}
	}()

	bufferRequest := make(chan int, 5)

	for k := 1; k < 6; k++ {
		bufferRequest <- k
	}
	close(bufferRequest)

	for t := range bufferRequest {
		<-bufferLimiter
		fmt.Println("Request ", t, " processed ", "at :", time.Now())
	}

}
