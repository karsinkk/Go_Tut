package main

import "fmt"

func main() {

	a := make([]string, 3)

	fmt.Println("init:", a)

	a[0] = "Karsin"
	a[1] = "Kamakotti"

	fmt.Println("set:", a)
	fmt.Println("size:", len(a))

	a = append(a, "Go")
	fmt.Println("apd:", a)

	b := make([]string, len(a))
	copy(b, a)
	fmt.Println("cpy:", b)

	b = append(b, "Programming", "Language")

	c := b[2:5]
	fmt.Println("slice:", c)
	d := b[2:]
	fmt.Println("slice:", d)
	e := b[:5]
	fmt.Println("slice:", e)

	f := []string{"a", "b", "c", "d", "e"}
	fmt.Println("dyn:", f)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("DynMDslice:", twoD)

}
