package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

func f1(arg int) (int, error) {
	if arg == 20 {
		return arg - 1, errors.New("can't work with 20")
	}
	return arg + 20, nil
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 20 {
		return -1, &argError{arg, "can't work with 20"}
	}
	return arg + 20, nil
}

func main() {
	if arg, err := f1(20); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arg)
	}

	if arg, err := f2(20); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arg)
	}

	_, e := f2(20)
	if arg, err := e.(*argError); err {
		fmt.Println(arg.arg)
		fmt.Println(arg.prob)
	}
}
