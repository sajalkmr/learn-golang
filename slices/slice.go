package main

import "fmt"

func main() {
	var intSlice = []int{3, 4, 5}

	intSlice = append(intSlice, 6, 7)
	fmt.Println(intSlice)

	nums := make([]int, 5)

	fmt.Println(nums)

}
