package main

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday, the start of the week.")
	case "Tuesday":
		fmt.Println("It's Tuesday, still early in the week.")
	case "Wednesday", "Thursday":
		fmt.Println("It's midweek, keep going!")
	case "Friday":
		fmt.Println("It's Friday, time to relax.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend, enjoy!")
	default:
		fmt.Println("Invalid day provided.")
	}
}
