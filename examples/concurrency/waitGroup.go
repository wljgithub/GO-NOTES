package main

import (
	"fmt"
	"sync"
	"time"
)

func expansiveTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("goroutine %d is beginning\n", id)
	// some expansive task
	time.Sleep(time.Second)
	fmt.Printf("goroutine %d has done\n", id)
}
func main() {
	var wg = new(sync.WaitGroup)
	var goRoutineNo = 5
	for i := 0; i < goRoutineNo; i++ {
		wg.Add(1)
		go expansiveTask(i, wg)
	}
	wg.Wait()
}
