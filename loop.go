package main

func main() {
	sum := 0
	names := []string{"홍길동", "이순신", "강감찬"}

	for i := 0; i <= 100; i++ {
		sum += i
	}

	println(sum)

	for index, name := range names {
		println(index, name)
	}
	variadic("This", "is", "go", "language")
}

func variadic(str ...string) {
	for _, c := range str {
		println(c)
	}
}
