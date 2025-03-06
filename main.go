package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// ONE WAY
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}

	// SECOND WAY
	//alex := person{"Alex", "Anderson"}

	fmt.Println(alex)
}
