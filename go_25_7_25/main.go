package main

func main() {
	a, b := saySomething("Hello")
	println(a, b)
}

func saySomething(s string) (string, string) {
	return s, "World"

}