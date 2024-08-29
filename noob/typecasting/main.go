package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int = 5
	fmt.Printf("value : %d ", num)
	fmt.Printf(", Type is %T\n", num)

	var data float64 = float64(num)

	data = data + 1.90

	fmt.Printf(" value of data flaot : %f ", data)
	fmt.Printf(", type of data float : %T\n", data)

	var d1 int = int(data)

	fmt.Printf(" value of data int : %d ", d1)
	fmt.Printf(", type of data int : %T\n", d1)

	fmt.Println("--------------------------------------------------------")
	fmt.Println()

	v := 8
	d := 3
	var r float64 = float64(v) / float64(d)

	fmt.Println("value of float r :", r)
	r = float64(int(r))
	fmt.Println("value of int r :", r)

	fmt.Println("--------------------------------------------------------")
	fmt.Println()

	var str string = "lucy"

	fmt.Println("value of  str :", str)
	n1 := 123
	str = strconv.Itoa(n1) // integer to alphabet , no need of error bca eveything is string

	fmt.Println("value of  str :", str)
	fmt.Printf("value of str : %T", str)

	fmt.Println("--------------------------------------------------------")
	fmt.Println()

	str1 := "1234"
	num1, error := strconv.Atoi(str1) // string to int require error message aswell to ignore use '_' insted of error

	fmt.Println(num1, ": message: ", error)

}
