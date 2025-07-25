package main

import "log"

type User struct {
	FirstName string
	LastName string
}

func main() {
	myMap := make(map[string]User)

	myMap["user1"] = User{FirstName: "Umesh", LastName: "Kedimi"}
	myMap["user2"] = User{FirstName: "Sujatha", LastName: "Kedimi"}

	log.Println("User 1:", myMap["user1"].FirstName, myMap["user1"].LastName)
	log.Println("User 2:", myMap["user2"].FirstName, myMap["user2"].LastName)

	me := User{
		FirstName: "Umesh", 
		LastName: "Kedimi",
	}

	myMap["me"] = me
	log.Println("Me:", myMap["me"].FirstName, myMap["me"].LastName)
}