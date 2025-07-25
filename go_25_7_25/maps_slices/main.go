package main

import "log"

func main() {
	myMap := make(map[string]int)

	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 15
	myMap["mango"] = 25

	log.Println("Initial map:", myMap)

	for key, value := range myMap {
	log.Printf("Key: %s, Value: %d\n", key, value)
	}
}