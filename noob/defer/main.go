package main

import "fmt"

func sum(a, b int) int {
	return a + b
}
func main() {
	ans := sum(2, 3)
	ans = ans * 2
	defer fmt.Println("first defer: ", ans)
	defer fmt.Println("second defer :", ans)
	fmt.Println("end of line")

}
