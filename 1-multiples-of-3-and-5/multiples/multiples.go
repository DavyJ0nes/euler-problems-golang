package multiples

// Multiples describes a slice of ints
type Multiples []int

// NewOf3Or5 gives a slice of numbers below
// the desried maximum that are multiples of
// either 3 or 5
func NewOf3Or5(max int) Multiples {
	var m Multiples

	for i := 1; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			m = append(m, i)
		}
	}

	return m
}

// Sum simply sums multiples together
func (m Multiples) Sum() int {
	var output int
	for _, i := range m {
		output += i
	}

	return output
}
