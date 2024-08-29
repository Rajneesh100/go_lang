package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "ram mohan atul vikash"
	list := strings.Split(data, ",")
	fmt.Println(list)

	data = "oneone two three four"
	count := strings.Count(data, "one")
	fmt.Println(count)

	data = "      HI anshu !        "
	data = strings.TrimSpace(data)
	fmt.Println(data)
	var s1 string = "Rajneesh"
	var s2 string = "kumar"

	s1 = strings.Join([]string{s1, "kumar", s2}, ",,")
	fmt.Println(s1)
}
