package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "sam"
	name[4] = "ram"
	fmt.Printf("names of people : %q\n", name)

	var number = [6]int{1, 2, 3, 4}
	fmt.Println(number)

	fmt.Printf("size of number :%s, %d\n", name[0], len(number))

}
