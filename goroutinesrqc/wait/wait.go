package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Printf("Goroutine test %d\n", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
