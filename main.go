package main

import "fmt"

func hello() {
	fmt.Println("Hello, World!")
}

func dataTypes() {
	var a int = 10       // int8, int16, int32, int64
	var b float64 = 3.14 // float32
	var c string = "Hello, World!"
	var d bool = true

	fmt.Println(a, b, c, d)
}

func addNumbers(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(addNumbers(42, 13))
}
