package main

import "fmt"

func devide(a, b float64) (float64, error) {
	if b == 0 && a > 0 {
		return a / b, fmt.Errorf("devision by zero")
	} else {
		return a / b, nil
	}
}

func main() {
	fmt.Println(devide(10, 0))
	ans, _ := devide(10, 9)
	fmt.Printf("ans: %.3f\n", ans)

}
