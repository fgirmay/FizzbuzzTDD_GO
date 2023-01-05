package main

import (
	"log"
	"strconv"
)

func Fizzbuzz(input string) (output string) {
	inputNum, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("%v is not a number", input)
		return input
	}

	if inputNum % 3 == 0 {
		output += "fizz"
	}
	if (inputNum % 5 == 0) {
		output += "buzz" 
	}
	return output
}