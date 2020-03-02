package main

import (
	"fmt"
	"time"
)

func expansiveTask(id int, ch chan int, task int) {
	fmt.Printf("I am go routine no:%d,i am processing task %d\n", id, task)
	time.Sleep(10 * time.Second)
	ch <- id

}
func main() {
	var goRoutineNo = 10
	var scheduler = make(chan int, goRoutineNo)
	var tasks = make(chan int, goRoutineNo)

	go func() {
		for i := 0; true; i++ {
			tasks <- i
		}
	}()
	for i := 0; i < goRoutineNo; i++ {
		task := <-tasks
		go expansiveTask(i, scheduler, task)
	}

	for {
		select {
		case goroutineNo := <-scheduler:
			task := <-tasks
			go expansiveTask(goroutineNo, scheduler, task)
		case <-time.After(2 * time.Second):
			fmt.Printf("goroutine is busy,and has stuck 2 for second\n")
		}
	}
}
