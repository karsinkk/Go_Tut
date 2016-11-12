package main

import (
	"fmt"
	"os"
)

func main() {

	f := CreateFile("file.txt")
	defer CloseFile(f)
	WriteFile(f)

}

func CreateFile(p string) *os.File {
	fmt.Println("Creating File")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func WriteFile(f *os.File) {
	fmt.Println("Writing to File")
	fmt.Fprintln(f, "Karsin Kamakotti")
}

func CloseFile(f *os.File) {
	fmt.Println("Closing File")
	f.Close()
}
