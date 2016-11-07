package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "Karsin Kamakotti!!" }()

	msg := <-messages
	fmt.Println(msg)

}
