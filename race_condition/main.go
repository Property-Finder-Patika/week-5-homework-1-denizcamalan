package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
    counter int
	  // used to wait for the program to finish
   	wg sync.WaitGroup 
)

func main() {
	wg.Add(2)
// Create two goroutines.
	go goCounter(3)
	go goCounter(5)

// Wait for the goroutines to finish.
wg.Wait()
fmt.Println("Final Count:", counter)
}

func goCounter(num int) {
	
for i := 1; i <= num; i++ {
	
		temp := counter
		runtime.Gosched()
		temp++
		counter = temp
		
}
// call to Done to tell main we are done.
wg.Done()
}
