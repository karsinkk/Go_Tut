package main

import "fmt"

func newFunc() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func newPrint() func() {
	i := 0
	return func() {
		i += 1
		fmt.Println("Iter: ", i)
	}
}

func main() {
	Func := newFunc()
	fmt.Println("Iter 1:", Func())
	fmt.Println("Iter 2:", Func())
	fmt.Println("Iter 3:", Func())

	Func1 := newFunc()
	fmt.Println("Init:", Func1())

	Func2 := newPrint()
	Func2()
	Func2()
	Func2()

}
