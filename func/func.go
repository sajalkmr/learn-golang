package main

import "fmt"

func addNumbers(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(addNumbers(17, 10))
}
