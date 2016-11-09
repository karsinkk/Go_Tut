package main

import "fmt"

func main() {

	chan1 := make(chan string, 2)

	chan1 <- "Karsin"
	chan1 <- "Kamakotti"

	close(chan1)

	for i := range chan1 {
		fmt.Println(i)
	}

}
