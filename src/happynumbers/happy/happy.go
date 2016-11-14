package happy

func IsHappy(x int) bool {
	n := SumSquaredDigits(x)

	switch n {
	case 1:
		return true
	case 4:
		return false
	default:
		return IsHappy(n)
	}
}

func SumSquaredDigits(n int) int {
	return IntSum(IntSquareSeq(IntPopIter(n)))
}

func IntSum(xs chan int) int {
	x := 0

	for i := range xs {
		x += i
	}

	return x
}

func IntSquareSeq(xs chan int) chan int {
	c := make(chan int)

	go func() {
		for i := range xs {
			c <- i * i
		}

		close(c)
	}()

	return c
}

func IntPopIter(n int) chan int {
	c := make(chan int)

	go func() {
		var i int

		for n > 0 {
			i, n = IntPop(n)
			c <- i
		}

		close(c)
	}()

	return c
}

func IntPop(n int) (head, tail int) {
	head = n % 10
	tail = n / 10
	return
}
