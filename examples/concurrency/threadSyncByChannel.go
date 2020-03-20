package main

import (
	"fmt"
	"time"
)

func expansiveTask(id int, done chan bool) {
	// simulate expansive task
	time.Sleep(time.Second)
	fmt.Printf("Task %d has done\n", id)
	done <- true
}
func main() {
	ch := make(chan bool)
	taskNo := 10

	for i := 0; i < taskNo; i++ {
		go expansiveTask(i, ch)
	}
	for i := 0; i < taskNo; i++ {
		<-ch
	}
}
