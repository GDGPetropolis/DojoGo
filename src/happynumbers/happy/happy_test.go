package happy

import "testing"

func TestIsHappy(t *testing.T) {
	happyNumbers := []int{1, 7, 10, 13, 19, 23, 28, 31, 32, 44, 49, 68, 70,
		79, 82, 86, 91, 94, 97, 100, 103, 109, 129, 130, 133, 139, 167, 176, 188,
		190, 192, 193, 203, 208, 219, 226, 230, 236, 239, 262, 263, 280, 291, 293,
		301, 302, 310, 313, 319, 320, 326, 329, 331, 338, 356, 362, 365, 367, 368,
		376, 379, 383, 386, 391, 392, 397, 404, 409, 440, 446, 464, 469, 478, 487,
		490, 496, 536, 556, 563, 565, 566, 608, 617, 622, 623, 632, 635, 637, 638,
		644, 649, 653, 655, 656, 665, 671, 673, 680, 683, 694, 700, 709, 716, 736,
		739, 748, 761, 763, 784, 790, 793, 802, 806, 818, 820, 833, 836, 847, 860,
		863, 874, 881, 888, 899, 901, 904, 907, 910, 912, 913, 921, 923, 931, 932,
		937, 940, 946, 964, 970, 973, 989, 998, 1000}

	for i, happy := range happyNumbers {
		if !IsHappy(happy) {
			t.Errorf("isHappy(%d) == %t, want %t", happy, false, true)
		}

		if happy == 1000 {
			break
		}

		for sad := happy + 1; sad < happyNumbers[i+1]; sad++ {
			if IsHappy(sad) {
				t.Errorf("IsHappy(%d) == %t, want %t", sad, true, false)
			}
		}
	}
}

func TestSumSquqredDigits(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1, 1},
		{0, 0},
		{21, 5},
		{12, 5},
		{12321, 19},
	}

	for _, c := range cases {
		if got := SumSquaredDigits(c.in); got != c.want {
			t.Errorf("SumSquaredDigits(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestIntSquareSeq(t *testing.T) {
	c := make(chan int)
	start, stop := 1, 1000

	go func() {
		for i := start; i < stop; i++ {
			c <- i
		}

		close(c)
	}()

	i := start

	for got := range IntSquareSeq(c) {
		want := i * i

		if got != want {
			t.Errorf("IntSquareSeq([%d,%d)) yields %d, want %d", start, stop, got,
				want)
		}

		i++
	}
}

func TestIntPopIter(t *testing.T) {
	cases := []struct {
		in   int
		want []int
	}{
		{1, []int{1}},
		{1234, []int{4, 3, 2, 1}},
		{4321, []int{1, 2, 3, 4}},
		{1010, []int{0, 1, 0, 1}},
	}

	for _, c := range cases {
		i := 0
		for got := range IntPopIter(c.in) {
			if i >= len(c.want) {
				t.Errorf("# of IntPopIter(%d) greater than expected, want %d", c.in,
					len(c.want))
			}

			if got != c.want[i] {
				t.Errorf("IntPopIter(%d) yields %d, want %d", c.in, got, c.want[i])
			}

			i++
		}

		if i < len(c.want) {
			t.Errorf("# of IntPopIter(%d) smaller than expected, want %d", c.in,
				len(c.want))
		}
	}
}

func TestIntPop(t *testing.T) {
	cases := []struct {
		in, head, tail int
	}{
		{7, 7, 0},
		{2314, 4, 231},
		{1510, 0, 151},
		{151, 1, 15},
		{15, 5, 1},
		{1, 1, 0},
		{0, 0, 0},
	}

	for _, c := range cases {
		head, tail := IntPop(c.in)
		if head != c.head || tail != c.tail {
			t.Errorf("IntPop(%d) = (%d, %d), want (%d, %d)", c.in, head, tail,
				c.head, c.tail)
		}
	}
}
