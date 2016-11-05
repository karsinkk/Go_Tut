package main

import "fmt"

type data struct {
	name string
	id   int
}

func main() {
	fmt.Println(data{"Karsin", 21})
	fmt.Println(data{name: "Prina", id: 28})
	fmt.Println(data{name: "Gany"})
	fmt.Println(&data{name: "Kamakotti", id: 843})

	a := data{"Karsin Kamakotti", 23}
	fmt.Println("Name:", a.name, " | ID:", a.id)

	a.id = 67
	fmt.Println("Name:", a.name, " | New ID:", a.id)

	b := &a
	b.id = 84
	fmt.Println("Name:", b.name, " | New ID:", b.id)

}
