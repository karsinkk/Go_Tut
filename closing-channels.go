package main

import "fmt"

var done = make(chan bool, 1)

func printshit(chan1 chan string) {
	for {
		i, more := <-chan1
		if more {
			fmt.Println(i)
		} else {
			fmt.Println("Got all Shit!!")
			done <- true
			return
		}
	}
}

func main() {

	chan1 := make(chan string)

	go printshit(chan1)

	for i := 0; i < 5; i++ {
		chan1 <- fmt.Sprintf("Karsin :%d", i)
	}

	close(chan1)
	fmt.Println("Done Shit!!")
	<-done

}
