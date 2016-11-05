package main

import "fmt"

type data struct {
	width, height int
}

func (d *data) area() int {
	return d.width * d.height
}

func (d data) peri() int {
	return 2*d.width + 2*d.height
}

func main() {

	a := data{32, 54}

	fmt.Println("Area:", a.area())
	fmt.Println("Perimeter:", a.peri())

	b := &a
	fmt.Println("Area:", b.area())
	fmt.Println("Perimeter:", b.peri())

}
