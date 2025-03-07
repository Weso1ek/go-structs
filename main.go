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

	// ******** STRUCTS **********************************************************

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

	// ******** MAPS **********************************************************

	// declare 1 method
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#00ff00",
		"green": "#00ff00",
	}

	// declare 2 method

	var colors2 map[string]string // zero value empty map

	// declare 3 method

	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"

	delete(colors3, "white") // delete map element

	//fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	printMap(colors)

	// ******** MAPS **********************************************************
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func getMonthlyPrice(tier string) int {
	if tier == "basic" {
		return 10000
	} else if tier == "premium" {
		return 15000
	} else if tier == "enterprise" {
		return 50000
	}

	return 0
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
