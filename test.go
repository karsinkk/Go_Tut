package main

import "fmt"

func main() {

	done := make(chan bool)

	go func(arg int) {
		for i := 0; i < arg; i++ {
			fmt.Println("Karsin")
		}
		done <- true
	}(11)

	fmt.Println("Kamakotti")
	<-done

}
