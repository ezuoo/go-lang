package main

import (
	"fmt"
)

func main() {

	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	var arr2 = [3]int{4, 5, 6}
	var arr3 = [...]int{7, 8, 9}

	fmt.Println(arr, arr2, arr3)

	var aarr = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(aarr)
}
