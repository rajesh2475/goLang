package main

import "fmt"

type address struct {
	address1 string
	address2 string
	zipCode  int
}

type person struct {
	firstName string
	lastName  string
	phoneNum  string
	age       int
	address   // address => address(vName) address(Type)
}

func main() {
	// Method 1
	// details := person{"Rajesh", "M", "7760832839", 26}

	// Method 2
	details := person{
		firstName: "Rajesh",
		lastName:  "M",
		phoneNum:  "7760832839",
		age:       26,
		address: address{
			address1: "Banaswadi",
			address2: "bangalore",
			zipCode:  560033,
		},
	}

	// Method 3
	// var details person
	// details.age = 26
	// details.firstName = "Rajesh"
	// details.lastName = "M"
	// details.phoneNum = "7760832839"
	// pointerAddress := &details
	// pointerAddress.update("Deepu")
	details.update("Deepu")
	details.printDetails()

}

func (p person) printDetails() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *person) update(firstName string) {
	p.firstName = firstName
}
