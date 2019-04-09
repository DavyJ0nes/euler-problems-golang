package main

import (
	"fmt"

	"github.com/davyj0nes/euler-problems-golang/1-multiples-of-3-and-5/multiples"
)

func main() {
	max := 1000
	m := multiples.NewOf3Or5(max)
	result := m.Sum()

	fmt.Printf(
		"The sum of all the multiples of 3 or 5 below %d = %d\n", max, result,
	)
}
