package main

import "fmt"

type Rs struct {
	Nth    int
	Result int
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	var a, b = 0, 1
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}
	return a
}
func worker(jobs <-chan int, result chan<- Rs) {
	for n := range jobs {
		rs := fib(n)
		result <- Rs{n, rs}
	}
}
func main() {

	jobs := make(chan int, 100)
	result := make(chan Rs, 100)
	go func() {
		for i := 0; true; i++ {
			jobs <- i
		}
	}()

	var goRoutineNo = 10
	go func() {
		for i := 0; i < goRoutineNo; i++ {
			go worker(jobs, result)
		}
	}()
	for {
		select {
		case rs := <-result:
			fmt.Printf("%d number of fibonnaci is:%d\n", rs.Nth, rs.Result)
		}
	}
}
