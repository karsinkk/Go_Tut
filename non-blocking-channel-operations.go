package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "Go"
		chan1 <- "Golang"
	}()

	select {
	case msg1 := <-chan1:
		fmt.Println("Blocking", msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("Non Blocking Receive timeout")
	default:
		fmt.Println("Non Blocking Receive")
	}

	msg2 := "Karsin"
	select {
	case chan1 <- msg2:
		fmt.Println("Blocking Send Karsin")
	default:
		fmt.Println("Non Blocking Send Karsin")
	}

	select {
	case chan1 <- "Kamakotti":
		fmt.Println("Blocking Send Kamakotti")
	default:
		fmt.Println("Non Blocking Send Kamakotti")
	}

	select {
	case msg3 := <-chan1:
		fmt.Println("Blocking Receive msg3", msg3)
	case msg4 := <-chan2:
		fmt.Println("Blocking Receive msg4", msg4)
	default:
		fmt.Println("Multiple Case")
	}

}
