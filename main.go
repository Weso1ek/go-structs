package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo // TIP można to zrobić inaczej bez nazwy pola
}

func main() {
	// ONE WAY
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
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
	alex3.contactInfo.email = "tescior@gmail.com"
	alex3.contactInfo.zipCode = 222

	//fmt.Println(alex)
	//fmt.Println(alex2)
	//fmt.Println(alex3)

	//alexPointer := &alex
	//alexPointer.updateName("Robocop")

	alex.updateName("Robocop")

	alex.print()
	alex2.print()
	alex3.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
