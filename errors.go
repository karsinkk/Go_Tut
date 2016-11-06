package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 28 {
		return -1, errors.New("Can't parse arg")
	}
	return arg + 3, nil
}

type argError struct {
	val int
	msg string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.val, e.msg)
}

func f2(arg int) (int, error) {
	if arg == 28 {
		return -1, argError{arg, "Can't be parsed"}
	}
	return arg + 4, nil
}

func main() {

	s := []int{12, 34, 64, 28, 96, 34}

	for _, data := range s {
		if v, msg := f1(data); msg == nil {
			fmt.Println("F1 parsed val!! : ", v)
		} else {
			fmt.Println(msg)
		}
	}

	for _, data := range s {
		if v, msg := f2(data); msg == nil {
			fmt.Println("F2 parsed val!! :", v)
		} else {
			fmt.Println(msg)
		}
	}

}
