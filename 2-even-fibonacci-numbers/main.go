package main

import (
	"fmt"

	"github.com/davyj0nes/euler-problems-golang/2-even-fibonacci-numbers/fibo"
)

func main() {
	max := 4000000
	result := fibo.SumEven(max)

	fmt.Printf(
		"The sum of all the even fibonacci numbers less than %d = %d\n", max, result,
	)
}
