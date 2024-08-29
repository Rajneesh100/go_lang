package main

import (
	"encoding/json"
	"fmt"
)

// import "fmt"

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("learning json")

	var person Person = Person{Name: "anshu", Age: 34, IsAdult: true}
	fmt.Println("person data is :", person)
	// fmt.Println("nmae :",person.name)
	//convert into json

	jsonData, er := json.Marshal(person)

	if er != nil {
		fmt.Println("got issue in conversion :", er)
		return
	}
	fmt.Println("after marshling :", string(jsonData))

	// unmarshling
	var personData Person
	er = json.Unmarshal(jsonData, &personData)
	if er != nil {
		fmt.Println("got some error in unmarshling: ", er)
	}

	fmt.Println("person data after unmarshling :", personData)

}
