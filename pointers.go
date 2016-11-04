package main

import "fmt"

func zeroval(val int) {
	val = 0
}

func zeroptr(val1 *int) {
	*val1 = 0
}

func main() {

	i := 1

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("addr:", &i)

}
