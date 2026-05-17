package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Println("Worker", id, "started")
	// Do some work here
	fmt.Println("Worker", id, "done")
	defer wg.Done() // Decrement the WaitGroup counter
}

//wg.Done() decrements the WaitGroup counter
//wg.Wait() blocks until the WaitGroup counter is zero	
//wg.Add(1) increments the WaitGroup counter

func main() {
	fmt.Println("Learning Synchronization in Go")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i,&wg)
	}
	wg.Wait() // Wait until the WaitGroup counter is zero
	fmt.Println("All workers done")
}
