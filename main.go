package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// ONE WAY
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex.anderson@gmail.com",
			zipCode: 9400,
		},
	}

	// SECOND WAY
	alex2 := person{"Robert", "Hawks", contactInfo{"seba.wesol@gmail.com", 9999}}

	// THIRD WAY
	var alex3 person

	alex3.firstName = "Simon"
	alex3.lastName = "Henderson"
	alex3.contact.email = "tescior@gmail.com"
	alex3.contact.zipCode = 222

	fmt.Println(alex)
	fmt.Println(alex2)
	fmt.Println(alex3)

	fmt.Printf("%+v\n", alex3)
}
