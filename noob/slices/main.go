package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5} // dynamic array like vector
	number = append(number, 10, 11, 12, 13, 14, 15, 16)
	fmt.Println(number)
	fmt.Printf("%T\n", number)

	name := []string{}
	fmt.Println(name)

	fmt.Println("slice: ", number)
	fmt.Println("len :", len(number))
	fmt.Println("cap :", cap(number))

	//            type, size, max capacity
	vector := make([]int, 3, 5)

	fmt.Println("slice: ", vector)
	fmt.Println("len :", len(vector))
	fmt.Println("cap :", cap(vector))

	vector = append(vector, 3, 4)

	fmt.Println("slice: ", vector)
	fmt.Println("len :", len(vector))
	fmt.Println("cap :", cap(vector))

	vector = append(vector, 9) // when size exceed it's capacity then capavity becomes double

	fmt.Println("slice: ", vector)
	fmt.Println("len :", len(vector))
	fmt.Println("cap :", cap(vector))

}
