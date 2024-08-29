package main

import "fmt"

func main() {
	n := 8
	// var ptr *int = &n
	// ptr = &n
	ptr := &n

	// adress of n : value
	fmt.Println(&n, ": ", n)

	// adress of n : value at adress
	fmt.Println(ptr, ": ", *ptr)

	// adress of pointer :  value of pointer (which is equal to adress of n)
	fmt.Println(&(ptr), ": ", ptr)

	var ptr1 *int

	fmt.Println(ptr1) // => <nil>
}
