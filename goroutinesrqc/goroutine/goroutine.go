package main

import (
	"fmt"
	"time"
)

func action() {
	fmt.Println("Test Goroutine")
}
func main() {
	go action()
	time.Sleep(2 * time.Second)
}
