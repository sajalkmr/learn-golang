package main

import (
	"fmt"
)

func main() {
	// Example 1: Converting int to float64
	var x int = 10
	y := float64(x)
	fmt.Printf("x: %d, y: %f\n", x, y)

	// Example 2: Converting float64 to int
	var a float64 = 5.7

	b := int(a)
	fmt.Printf("a: %f, b: %d\n", a, b)

	// Example 3: Converting int to string
	var num int = 42
	str := fmt.Sprintf("%d", num)
	fmt.Printf("num: %d, str: %s\n", num, str)

	var floatNum32 float32 = 3.14
	var intNum32 int32 = 3
	result := floatNum32 + intNum32
	fmt.Println(result)

}
