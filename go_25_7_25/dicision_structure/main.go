package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if !isTrue {
		log.Println("The condition is true")
	} else {
		log.Println("The condition is false")
	}

	count := 4

	if count > 5 {
		log.Println("Count is greater than 5")
	} else if count == 5 {
		log.Println("Count is equal to 5")
	} else {
		log.Println("Count is less than 5")
	}

	// Switch statement example
	switch count {
	case 1:
		log.Println("The condition when the case is 1")

	case 2:
		log.Println("The condition when the case is 2")

	case 6:
		log.Println("The condition when the case is 6")

	default:
		log.Println("The condition when the case is ALL")
	}
}
