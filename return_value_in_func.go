package main

func main() {
	// total := sum(1, 7, 3, 5, 9)

	count, total := return_way_2(1, 2, 3, 4, 5)
	println(count, total)
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func return_way_1(nums ...int) (int, int) {
	s := 0
	count := 0

	for _, i := range nums {
		s += i
		count++
	}
	return count, s

}

func return_way_2(nums ...int) (count int, total int) {

	for _, i := range nums {
		total += i
	}

	count = len(nums)
	return

}
