package main

import "fmt"

func main() {
	var intArr [5]int

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	intArr[0] = 34
	fmt.Println(intArr[0])

	fmt.Println(len(intArr))

	intArr2 := [3]string{"a", "b", "c"}
	
	fmt.Println(intArr2)

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])

	fmt.Println(&intArr2[0])

}
