package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func dif(a, b int) int {
	return a - b
}

func main() {

	res := sum(5, 98)
	fmt.Println("Sum: ", res)

	fmt.Println("Dif: ", dif(43, 54))

}
