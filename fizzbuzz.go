package main

import (
	"errors"
	"fmt"
)

//const qwest = 30

func div(i, div int) bool {
	if i%div == 0 {
		return true
	}
	return false
}

func input() (int, error) {
	fmt.Print("Enter text: ")
	var inputValue int
	fmt.Scanln(&inputValue)
	if inputValue == 0 {
		return inputValue, errors.New("foobar unexpected error")

	} else {
		return inputValue, nil
	}
}

func main() {
	var qwest, err = input();
	if err == nil {
		for i := 1; i <= qwest; i++ {
			if div(i, 3) == true && div(i, 5) == true {
				fmt.Println("fizzbuzz")
			} else if div(i, 3) == true {
				fmt.Println("fizz")
			} else if div(i, 5) == true {
				fmt.Println("buzz")
			}
		}
	} else {
		fmt.Print(err)
	}
}
