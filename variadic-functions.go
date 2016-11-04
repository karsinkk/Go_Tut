package main

import "fmt"

func calc(nums ...int) (int, int) {
	sum, prod := 0, 1
	for _, val := range nums {
		sum += val
		prod *= val
	}
	return sum, prod
}

func main() {
	sum, prod := calc(1, 3, 42, 56, 7, 3, 2, 5, 6, 7, 85)
	fmt.Printf("Sum:%d | Prod:%d\n", sum, prod)

	sli := []int{23, 43, 11, 4, 2, 5, 6, 7, 23}
	sum, prod = calc(sli...)
	fmt.Printf("Sum:%d | Prod:%d\n", sum, prod)

}
