package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to:", myString)
	changeUsingPointer(&myString)
	log.Println("After myString is now set to:", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Blue"
	*s = newValue
}