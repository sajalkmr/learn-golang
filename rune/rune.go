package main

import "fmt"

func main() {

	var r rune = '日'

	fmt.Printf("Rune: %c\n", r)
	fmt.Printf("Unicode code point: %U\n", r)

	codePoint := int(r)
	fmt.Printf("Code point as integer: %d\n", codePoint)

	str := string(r)
	fmt.Printf("String from rune: %s\n", str)
}
