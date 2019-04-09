package fibo

func SumEven(max int) int {
	var result int
	f := fibonacci()

	for {
		val := f()
		if val > max {
			return result
		}

		if val%2 == 0 {
			result += val
		}
	}
}

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}
