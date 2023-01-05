package main

import (
	"fmt"
	"strconv"
)

func Fizzbuzz(input string) (output string) {
	inputNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("%v is not a number\n", input)
		return input
	}

	if inputNum % 3 == 0 {
		output += "fizz"
	}
	if (inputNum % 5 == 0) {
		output += "buzz" 
	}

	if output == "" {
		output = "not fizzy or buzzy"
	}
	return
}