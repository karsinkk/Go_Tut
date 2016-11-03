package main

import "fmt"

func main() {

	a := []int{2, 3, 5}

	for i, j := range a {
		fmt.Printf("%d -> %d\n", i, j)
	}

	sum := 0

	for _, c := range a {
		sum += c
	}
	fmt.Println("Sum: ", sum)

	m := map[string]int{"k1": 34, "k2": 9238, "k3": 8498, "k4": -234}
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}

	for z, x := range "Karsin Kamakotti" {
		fmt.Printf("\"%d -> %c\" ", z, x)
	}
}
