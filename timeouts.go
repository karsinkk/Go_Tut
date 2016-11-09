package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		chan1 <- "Karsin"
	}()

	select {
	case msg1 := <-chan1:
		fmt.Println(msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout 1")
	}

	chan2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- "Kamakotti"
	}()

	select {
	case msg2 := <-chan2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout 3")
	}

}
