package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string
	mySlice = append(mySlice, "Sandeep")
	mySlice = append(mySlice, "Umesh")
	mySlice = append(mySlice, "Ram")

	log.Println("My Slice:", mySlice)
	log.Println("Length of My Slice:", len(mySlice))
	log.Println("Capacity of My Slice:", cap(mySlice))

	var mySlice2 = []int{2, 4, 5, 3}
	log.Println("My Slice 2:", mySlice2)

	sort.Ints(mySlice2)
	log.Println("Sorted My Slice 2:", mySlice2)

}
