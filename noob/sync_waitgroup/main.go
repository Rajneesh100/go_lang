package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d ended \n", i)
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// fmt.Println("workers ended")

	wg.Wait()
	fmt.Println("workers ended")

}
