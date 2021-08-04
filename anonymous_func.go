package main

import (
	"fmt"
)

func main() {
	sum := func(n ...int) int {
		v := 0
		for _, i := range n {
			v += i
		}

		return v
	}

	result := sum(1, 2, 3, 4, 5)
	fmt.Printf("합 : %d\n", result)
}
