package main

import "fmt"

func main() {

	if 6%3 == 0 {
		fmt.Println("6 is divisible by 3")
	} else {
		fmt.Println("6 is not divisible by 3")
	}

	if num := -2; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
