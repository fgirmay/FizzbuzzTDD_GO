package main

import "testing"

func TestFizzbuzz_Divisible_By_3(t *testing.T) {
	actual := Fizzbuzz("12")
	expected := "fizz"

	if actual != expected {
		t.Errorf("actual %s, expected %s", actual, expected)
	}
}

func TestFizzbuzz_Divisible_By_5(t *testing.T) {
	actual := Fizzbuzz("10")
	expected := "buzz"

	if actual != expected {
		t.Errorf("actual %s, expected %s", actual, expected)
	}
}
func TestFizzbuzz_Divisible_By_5_AND_3(t *testing.T) {
	actual := Fizzbuzz("15")
	expected := "fizzbuzz"

	if actual != expected {
		t.Errorf("actual %s, expected %s", actual, expected)
	}
}

func TestFizzbuzz_NOT_Divisible_By_5_AND_3(t *testing.T) {
	actual := Fizzbuzz("14")
	expected := "not fizzy or buzzy"

	if actual != expected {
		t.Errorf("actual %s, expected %s", actual, expected)
	}
}

func TestFizzbuzz_NOT_Number(t *testing.T) {
	actual := Fizzbuzz("abc")
	expected := "abc"

	if actual != expected {
		t.Errorf("actual %s, expected %s", actual, expected)
	}
}