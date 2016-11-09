package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 has expired")

	timer2 := time.NewTimer(time.Second * 1)

	go func() {
		<-timer2.C
		fmt.Printf("Timer 2 has expired")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 has been stopped")
	}
}
