package main

import (
	"fmt"
	"time"
)

func main() {
	const max_node = 500000000
	var node [max_node]bool

	startTime := time.Now()

	for i := 0; i < max_node; i++ {
		node[i] = true
	}
	// for i := (max_node / 2); i < max_node; i++ {
	// 	node[i] = true
	// }

	endTime := time.Since(startTime)

	fmt.Printf("실행시간: %f\n", endTime.Seconds())

	// for i := 0; i < 10; i++ {
	// 	println(node[i])
	// }
}
