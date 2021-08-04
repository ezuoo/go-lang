package main

func main() {

	avg := func(nums ...int) int {
		v := 0
		for _, num := range nums {
			v += num
		}
		return v / len(nums)
	}
	r1 := calc(avg, 1, 2, 3, 4, 5)
	println(r1)
}

func calc(f func(...int) int, n ...int) int {
	return f(n...)
}
