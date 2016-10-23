package dojo_test

import (
	"testing"
	"../dojo"
)

func TestFizzBuzzReturnFizz(t *testing.T) {
	var value = 3;
	var result = dojo.FizzBuzz(value);
	if result != "Fizz" {
		t.Error("Erro");
	}
}