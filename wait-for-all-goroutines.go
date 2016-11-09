package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printshit(arg int, str string) {
	defer wg.Done()
	for i := 0; i < arg; i++ {
		fmt.Println(str, i)
	}
}

func main() {

	wg.Add(2)
	go printshit(11, "Karsin")
	go printshit(6, "Kamakotti")

	wg.Wait()
}
