package main

import "testing"

func TestFizzbuzz_Divisible_By_3(t *testing.T) {
	actual := Fizzbuzz("12")
	expected := "fizz"

	if actual != expected {
		t.Errorf("actual %s, expected %s", actual, expected)
	}
}