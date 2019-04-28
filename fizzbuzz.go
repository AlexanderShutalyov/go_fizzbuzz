package main

import (
	"errors"
	"fmt"
)

func inputData() (int, error) {
	fmt.Print("Enter text: ")
	var inputValue int
	fmt.Scanln(&inputValue)
	if inputValue == 0 {
		return inputValue, errors.New("foobar unexpected error")
	} else {
		return inputValue, nil
	}
}


func division(i int) {

	if i%5 == 0 && i%3 == 0 {
		fmt.Println("fizzbuzz")
	} else if i%3 == 0 {
		fmt.Println("fizz")
	} else if i%5 == 0 {
		fmt.Println("buzz")
	}
}

func main() {
	if inputValue, err := inputData(); err == nil {
	for i := 1; i <= inputValue; i++ {
		division(i)
	}} else {
		fmt.Print(err)
	}
}

