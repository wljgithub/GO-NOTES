package main

import (
	"fmt"
	"os"
	"runtime"
	_ "syscall"
	"time"
	"net/http"
	_ "net/http/pprof"
)

func findDeath() {
	stats := new(runtime.MemStats)
	var bigSlice []string
	for {
		bigSlice := append(bigSlice, "abcdefg")
		fmt.Println(len(bigSlice), cap(bigSlice))

		runtime.ReadMemStats(stats)
		fmt.Printf("当前进程内存占用 %v MB", stats.Alloc/(1024*1024))

		time.Sleep(1e2)
	}
}

func HowManyGoRoutine() {
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println("i am goroutine", i)
		}()
	}
	fmt.Println(runtime.NumGoroutine())
}
func getPid() {
	fmt.Println(os.Getpid())
	fmt.Scanln()
}
func main() {


	 go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
	for i := 0; i < 10; i++ {
		findDeath()
	}
	fmt.Scanln()

}
