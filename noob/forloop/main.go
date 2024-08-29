package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i == 3 {
			continue
		}
		fmt.Println("number is :", i)

		if i == 4 {
			break
		}
	}

	number := []int{1, 2, 3, 4, 5} // dynamic array like vector
	number = append(number, 10, 11, 12, 13, 14, 15, 16)
	fmt.Println(number)
	fmt.Printf("%T\n", number)

	for index, value := range number {
		fmt.Println(index, ": ", value)
	}
}
