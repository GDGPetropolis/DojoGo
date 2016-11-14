package fizzbuzz

import "strconv"

func FizzBuzz(value int) string {
	var result = strconv.Itoa(value);

	if value % 3 == 0 && value % 5 == 0 {
		result = "Fizz Buzz";
	} else if value % 3 == 0 {
		result = "Fizz";
	} else if value % 5 == 0 {
		result = "Buzz";
	}

	return result;
}
