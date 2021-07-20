package main

import "fmt"

type address struct {
	address1 string
	address2 string
	pinCode  int
}

type ownerDetails struct {
	firstName string
	lastName  string
	age       int
	phoneNum  string
	address
}

type vehicle interface {
	getOwnerDetails() ownerDetails
}

type bike struct {
	noOfTyres int
	cost      int
}

type car struct {
	noOfTyres int
	cost      int
}

func main() {

	bike := bike{cost: 80000, noOfTyres: 2}
	car := car{cost: 1000000, noOfTyres: 4}

	getVehicleDetails(bike)
	getVehicleDetails(car)

}

func getVehicleDetails(v vehicle) {
	fmt.Printf("%+v", v)
	fmt.Println()
	fmt.Printf("%+v", v.getOwnerDetails())
	fmt.Println()
}

func (b bike) getOwnerDetails() ownerDetails {

	return ownerDetails{
		firstName: "Rajesh",
		lastName:  "M",
		age:       27,
		phoneNum:  "7760832839",
		address: address{
			address1: "banaswadi",
			address2: "Bangalore",
			pinCode:  560033,
		},
	}

}

func (b car) getOwnerDetails() ownerDetails {

	return ownerDetails{
		firstName: "Deepu",
		lastName:  "M",
		age:       26,
		phoneNum:  "7760832839",
		address: address{
			address1: "banaswadi",
			address2: "Bangalore",
			pinCode:  560033,
		},
	}

}
