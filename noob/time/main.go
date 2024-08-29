package main

import (
	"fmt"
	"time"
)

func main() {

	currtime := time.Now()
	fmt.Println(currtime)
	formatted := currtime.Format("2006/01/02 , Monday, 1:04 PM")
	fmt.Println(formatted)
	fmt.Printf("format of time %T\n", formatted)
}
