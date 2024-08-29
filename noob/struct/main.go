package main

import "fmt"

type person struct {
	name string
	age  int
}

type contact struct {
	email string
	phone string
}

type address struct {
	address string
}

type employee struct {
	person_detail  person
	contact_detail contact
	address_detail address
}

func main() {

	var p1 person

	fmt.Println(p1.name, "; ", p1.age)
	p1.name = "ram"
	p1.age = 34

	fmt.Println(p1.name, "; ", p1.age)

	var p2 = new(person)

	p2.name = "roma"
	p2.age = 22

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("\n-----------------------------------------------------\n")
	var e1 employee
	e1.person_detail.name = "dhruti"
	e1.person_detail.age = 22
	e1.contact_detail.email = "dhruti@gmail.com"
	e1.contact_detail.phone = "8875226534"
	e1.address_detail.address = "india, IIT Patna"

	fmt.Println(e1)

	e2 := employee{
		person_detail: person{
			name: "lizz",
			age:  90,
		},
		contact_detail: contact{
			email: "lizz@gmail.com",
			phone: "7009978344",
		},
		address_detail: address{
			address: "USA, montrial", // Add the missing comma here
		},
	}

	fmt.Println(e2)

}
