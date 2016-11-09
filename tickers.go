package main

import (
	"fmt"
	"time"
)

func main() {

	ticker1 := time.NewTicker(time.Millisecond * 600)

	go func() {
		for i := range ticker1.C {
			fmt.Println("Event at:", i)
		}
	}()

	time.Sleep(time.Millisecond * 2400)
	ticker1.Stop()
	fmt.Println("Ticker Stopped")

}
