package main

import "fmt"

func work(msg chan bool) {
	fmt.Println("In function")
	msg <- true
}

func main() {

	messages := make(chan bool)

	go work(messages)

	<-messages

}
