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

	go func(arg int) {
		for i := 0; i < arg; i++ {
			fmt.Println("Kamakotti")
		}
		done <- true
	}(6)

	fmt.Println("GO")
	<-done

}
