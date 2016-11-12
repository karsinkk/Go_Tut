package main

import "os"

func main() {
	panic(" Golang")

	_, err := os.Create("/tmpt/files")
	if err != nil {
		panic(err)
	}
}
