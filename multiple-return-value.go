package main

import "fmt"

func comb(a, b int) (int, int) {
	return a + b, a - b
}

func main() {

	a, b := comb(57, 34)
	fmt.Printf("Sum:%d | Dif:%d\n", a, b)

	_, d := comb(87482, 9238928)
	fmt.Println("Dif:", d)

}
