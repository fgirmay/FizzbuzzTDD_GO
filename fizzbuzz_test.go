package main

import (
	"fmt"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	t.Helper()
	cases := []struct{ Name, Input, Expected string} {
		{ "Divisible by 3", "12", "fizz" },
		{ "Divisible by 5", "10", "buzz" },
		{ "Divisible by 3 and 5", "15", "fizzbuzz" },
		{ "Not divisible by 5 and 3", "11", "not fizzy or buzzy" },
		{ "Not a number", "abc", "abc1" },
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("actual %s, expected %s", tc.Input, tc.Expected), func(t *testing.T) {
			actual := Fizzbuzz(tc.Input)
			if actual != tc.Expected {t.Fatal("Failure")}
		})
	}
}

// func TestFizzbuzz_Divisible_By_3(t *testing.T) {
// 	actual := Fizzbuzz("12")
// 	expected := "fizz"

// 	if actual != expected {
// 		t.Errorf("actual %s, expected %s", actual, expected)
// 	}
// }

// func TestFizzbuzz_Divisible_By_5(t *testing.T) {
// 	actual := Fizzbuzz("10")
// 	expected := "buzz"

// 	if actual != expected {
// 		t.Errorf("actual %s, expected %s", actual, expected)
// 	}
// }
// func TestFizzbuzz_Divisible_By_5_AND_3(t *testing.T) {
// 	actual := Fizzbuzz("15")
// 	expected := "fizzbuzz"

// 	if actual != expected {
// 		t.Errorf("actual %s, expected %s", actual, expected)
// 	}
// }

// func TestFizzbuzz_NOT_Divisible_By_5_AND_3(t *testing.T) {
// 	actual := Fizzbuzz("14")
// 	expected := "not fizzy or buzzy"

// 	if actual != expected {
// 		t.Errorf("actual %s, expected %s", actual, expected)
// 	}
// }

// func TestFizzbuzz_NOT_Number(t *testing.T) {
// 	actual := Fizzbuzz("abc")
// 	expected := "abc"

// 	if actual != expected {
// 		t.Errorf("actual %s, expected %s", actual, expected)
// 	}
// }