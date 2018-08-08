package main

import (
	"fmt"
	"sync"
)

var Count int = 0

func nonThreadSecurity(threadNums int) {
	for i := 0; i < threadNums; i++ {
		go func() {
			Count++
		}()
	}
	fmt.Scanln()
}

func threadSecurity(lock *sync.Mutex, threadNums int) {
	for i := 0; i < threadNums; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			Count++
		}()
	}
	fmt.Scanln()
}
func main() {
	var lock = &sync.Mutex{}
	threadNums := 9000
	// nonThreadSecurity(threadNums)
	threadSecurity(lock, threadNums)
	fmt.Println(Count)

}
