package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 2)
	pongs := make(chan string, 2)

	ping(pings, "Karsin")
	ping(pings, "Kamakotti")

	pong(pings, pongs)
	pong(pings, pongs)

	fmt.Println(<-pongs)
	fmt.Println(<-pongs)
}
