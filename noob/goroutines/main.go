package main

import (
	"fmt"
	"time"
)

func sayhello() {
	fmt.Println("sayhello:)")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("sayhello Done")
}

func sayhi() {
	fmt.Println("sayhi:)")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("sayhi Done")
}

func main() {
	go sayhello()
	go sayhi()

	time.Sleep(100 * time.Millisecond)

}
