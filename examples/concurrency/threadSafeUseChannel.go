package main

import (
	"fmt"
)

var count = 0

func plusOneUseChannel(result chan int) {
	result <- 1
}
func main() {

	var goRoutineNo = 1000

	// implement plus one use channel
	var result int
	ch := make(chan int)
	for i := 0; i < goRoutineNo; i++ {
		go plusOneUseChannel(ch)
	}
	for i := 0; i < goRoutineNo; i++ {
		r := <-ch
		result += r
	}
	fmt.Println(result)

}
