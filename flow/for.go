package main

import "fmt"


func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}



func main() {
	sum := 0
	i := 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}
