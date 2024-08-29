package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		fmt.Println(i)
	}
	marks := make(map[string]int)
	marks["anshu"] = 90
	marks["rajnii"] = 97
	fmt.Println(marks)
	for key, value := range marks {
		fmt.Println(key, ": ", value)
	}

	fmt.Println("marks of rajnii :", marks["rajnii"])
	marks["rajnii"] = 100

	fmt.Println("marks of rajnii :", marks["rajnii"])
	delete(marks, "anshu")
	for key, value := range marks {
		fmt.Println(key, ": ", value)
	}

	value, exists := marks["anshu"]

	fmt.Println("marks of anshu :", value)
	fmt.Println("anshu exists? :", exists)

	value1, exists1 := marks["rajnii"]

	fmt.Println("marks of anshu :", value1)
	fmt.Println("anshu exists? :", exists1)

	fmt.Println("-------------------------------------------")

	person := map[string]int{
		"roma": 80,
		"tess": 90,
	}

	person["lizz"] = 95

	for key, value := range person {
		fmt.Println(key, ": ", value)
	}

}
