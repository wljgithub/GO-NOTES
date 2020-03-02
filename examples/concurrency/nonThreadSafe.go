package main

import (
	"fmt"
	"sync"
)

var count = 0

func plusOne(wg *sync.WaitGroup) {
	defer wg.Done()
	count++
}
func main() {
	var goRoutineNo = 1000
	var wg = new(sync.WaitGroup)
	for i := 0; i < goRoutineNo; i++ {
		wg.Add(1)
		go plusOne(wg)
	}
	wg.Wait()
	// expect to 1000,but sometime is not
	fmt.Println(count)
}
