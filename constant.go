package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 2394879240

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(float64(d))

	fmt.Println(math.Sin(n))

}
