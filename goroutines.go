package main

import "fmt"

func printshit(str string, arg int) {
	for i := 0; i < arg; i++ {
		fmt.Printf("%s -> %d\n", str, i)
	}
}

func main() {

	printshit("WTF!!", 4)

	go printshit("Karsin", 5)
	go printshit("Kamakotti", 3)

	go func(str string) {
		fmt.Println("Printing this shit: ", str)
	}("Hello!")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
