package main

import "log"
import "time"

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	user := User{
		FirstName: "Umesh",
		LastName: "Kedimi",
		PhoneNumber: "1234567890",
		Age: 30,
		BirthDate: time.Date(1993, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	log.Println("User created:", user)
	log.Println(user.FirstName, user.LastName, "Phone:", user.PhoneNumber, "Age:", user.Age, "BirthDate:", user.BirthDate)
}

