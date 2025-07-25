package main

import "log"
// import "time"

type User struct {
	FirstName string
}

func (f *User) GetFullName() string {
	return f.FirstName + " " + "Kedimi"
}

func main() {
	var user1 User
	user1.FirstName = "Umesh"

	user2 := User{ FirstName: "Sujatha" }

	log.Println("User 1:", user1.GetFullName())
	log.Println("User 2:", user2.GetFullName())
}

