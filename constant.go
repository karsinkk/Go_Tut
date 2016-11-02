package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 3737983498394893

	const d = 3e23 / n
	fmt.Println(d)

	fmt.Println(float64(d))

	fmt.Println(math.Sin(n))

}
