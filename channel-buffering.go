package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "Karsin"
	messages <- "Kamakotti"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
