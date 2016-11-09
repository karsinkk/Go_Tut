package main

import (
	"fmt"
	"time"
)

func printshit(arg int, str string, done chan bool) {
	for i := 0; i < arg; i++ {
		fmt.Println(str, i)
	}
	done <- true
}

func check(done chan bool) {
	for range done {
		<-done
	}

}

func main() {

	done := make(chan bool)

	go printshit(11, "Karsin", done)
	go printshit(6, "Kamakotti", done)
	go check(done)

	time.Sleep(3 * 1e9)

}
