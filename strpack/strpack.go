package main

import (
	"fmt"
	"strings"
)

func main() {
	// Example 1: String Concatenation
	strSlice := []string{"Hello", "World", "Go"}
	joinedString := strings.Join(strSlice, ", ")
	fmt.Println("Joined string:", joinedString)

	// Example 2: Substring Search
	str := "Hello, World!"
	contains := strings.Contains(str, "World")
	fmt.Println("Does the string contain 'World'?", contains)

	index := strings.Index(str, "World")
	fmt.Println("Index of 'World' in the string:", index)

	// Example 3: String Replacement
	replacedString := strings.Replace(str, "World", "Go", -1)
	fmt.Println("Replaced string:", replacedString)

	// Example 4: Splitting Strings
	splitString := strings.Split(str, ", ")
	fmt.Println("Split string:", splitString)

	// Example 5: String Trimming
	trimmedString := strings.TrimSpace("   Hello, World!   ")
	fmt.Println("Trimmed string:", trimmedString)

	// Example 6: Case Conversion
	lowerString := strings.ToLower(str)
	upperString := strings.ToUpper(str)
	fmt.Println("Lowercase:", lowerString)
	fmt.Println("Uppercase:", upperString)

	// Example 7: Padding
	repeatedString := strings.Repeat("Go", 3)
	fmt.Println("Repeated string:", repeatedString)

	// Example 8: Checking Prefix and Suffix
	hasPrefix := strings.HasPrefix(str, "Hello")
	hasSuffix := strings.HasSuffix(str, "World!")
	fmt.Println("Does the string have prefix 'Hello'?", hasPrefix)
	fmt.Println("Does the string have suffix 'World!'?", hasSuffix)

	// Example 9: Counting Substrings
	count := strings.Count(str, "l")
	fmt.Println("Count of 'l' in the string:", count)
}
