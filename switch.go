package main

import (
	"fmt"
	"time"
)

func main() {

	i := 3
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!!")
	default:
		fmt.Println("It's a weekday....")
	}

	t := time.Now()
	fmt.Println(t)
}
