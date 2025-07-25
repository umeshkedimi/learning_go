package main

import "log"

type User struct {
	FirstName 	string
	LastName 	string
	Age 		int
}

func main() {
	// slice
	cities := []string{"Hyderabad", "Bengaluru", "Chennai", "Mumbai"}

	cities = append(cities, "Goa")

	for i, value := range cities {
		log.Println(i+1, "-", value)
	}

	// map
	animals := make(map[string]string)

	animals["cat"] = "Catty"
	animals["dog"] = "Doggy"
	animals["bird"] = "Duck"

	for key, value := range animals {
		log.Println(key, "name is", value)
	}


	// string
	firstLine := "This is for the first line"

	for i, v := range firstLine {
		log.Println(i, v)
	}

	// struct
	users := []User{}

	users = append(users, User{FirstName: "Umesh", LastName: "Kedimi", Age: 32})
	users = append(users, User{FirstName: "Sandeep", LastName: "K", Age: 32})

	for i, v := range users {
		log.Println(i+1,"-->",v.LastName, v.FirstName, v.Age)
	}


}
