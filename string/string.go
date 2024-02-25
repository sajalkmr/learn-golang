package main

import "fmt"

func main() {
	var myStr string = "Hello, folks!"

	var myStr2 string = `Hello, folks!`

	fmt.Println(myStr)
	fmt.Println(myStr2)
	fmt.Println()

}

func main() {

	var str string
	str = "Hello, World!"
	fmt.Println("Original string:", str)

	length := len(str)
	fmt.Println("Length of string:", length)

	char := str[0]
	fmt.Println("First character:", string(char))

	str1 := "Hello, "
	str2 := "Go!"
	newStr := str1 + str2
	fmt.Println("Concatenated string:", newStr)

	subStr := str[7:12]
	fmt.Println("Substring:", subStr)

}
