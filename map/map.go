package main

import "fmt"

func main() {

	mapEg := map[string]int{
		"Alex":    93,
		"Bob":     85,
		"Charlie": 72,
	}

	fmt.Println(mapEg)

	myMap := make(map[int]string)

	myMap[1] = "One"
	myMap[2] = "Two"
	myMap[3] = "Three"

	fmt.Println("Value associated with key 1:", myMap[1])

	delete(myMap, 1)

	fmt.Println("Map after deleting key 1:", myMap)

	_, exists := myMap[1]
	fmt.Println("Does key 1 exist:", exists)
}
