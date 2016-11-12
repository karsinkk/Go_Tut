package main

import (
	"fmt"
	"sort"
)

func main() {

	str := []string{"Karsin", "Kamakotti", "Go", "Programming"}
	sort.Strings(str)
	fmt.Println("Sorted Strings : ", str)

	ints := []int{1, 35, 35, 34, -3, 84, 310, -198391}
	sort.Ints(ints)
	fmt.Println("Ints Sorted : ", ints)

	check := sort.IntsAreSorted(ints)
	fmt.Println(check)

}
