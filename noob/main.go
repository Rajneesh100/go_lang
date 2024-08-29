package main

import (
	"fmt"
	"noob/function"
	"noob/myutil"
)

func main() {
	fmt.Println("hello there! ")
	myutil.Show("hi")

	var a int = 6
	fmt.Println(a)
	var ans rune = 'a' // during run time it will figur out that ans is a string
	ans = 'b'
	fmt.Println(ans)

	b := 8
	fmt.Println(b)
	fmt.Println(function.Solve(5))
	c := 9
	d := 56
	fmt.Println(function.Multi(c, d))
}
