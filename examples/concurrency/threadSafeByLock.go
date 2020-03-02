package main

import (
	"fmt"
	"sync"
)

var count int

func plusOneInThreadSafe(wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	count++
	lock.Unlock()
}
func main() {
	var goRounetineNo = 1000
	var wg = new(sync.WaitGroup)
	var lock = new(sync.Mutex)

	for i := 0; i < goRounetineNo; i++ {
		wg.Add(1)
		go plusOneInThreadSafe(wg, lock)
	}
	wg.Wait()
	fmt.Println(count)
}
