package main

import "log"

var s = "Seven"

func main() {
	var s2 = "Eight"

	log.Println("s is set to:", s)
	log.Println("s2 is set to:", s2)

	saySomething(s2)
}

func saySomething(s3 string) (string, string) {
	log.Println("saySomething called with:", s)
	return s3, "World"

}