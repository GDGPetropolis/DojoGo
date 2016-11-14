package fizzbuzz

import "testing"

func TestFizzBuzzReturnFizz(t *testing.T) {
	var value = 3;
	var result = FizzBuzz(value);
	if result != "Fizz" {
		t.Error("Erro");
	}
}
