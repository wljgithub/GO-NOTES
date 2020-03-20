package main

import (
	"fmt"
	"os"
	"time"
)

func concurrency(ch chan int, num int) {
	ch <- num
}
func forumChannel() {

	ch1, ch2 := make(chan int, 3), make(chan int)

	go func() {
		v, ok, s := 0, false, ""
		for {
			select {
			case v, ok = <-ch1:
				fmt.Println("read a")
				s = "a"
			case v, ok = <-ch2:
				fmt.Println("read b")
				s = "b"
			case <-time.After(100 * time.Millisecond):
				break
			}
			if ok {
				fmt.Println("v&s", v, s)
			} else {
				os.Exit(1)
			}
		}
	}()

	for i := 1; i < 4; i++ {
		select {
		case ch1 <- i:
			fmt.Println("in:",i)
		case ch2 <- i:
		}
	}

	fmt.Println("send over")

	//close(ch1)
	//close(ch2)

	fmt.Println("ch", ch1, ch2)
	select {} // 阻塞进程
}
func main() {
	//ch := make(chan int)
	//sum := 0
	//for i := 0; i < 10000000; i++ {
	//	go concurrency(ch, i)
	//}
	//go func() {
	//	for {
	//		i := <-ch
	//		sum += i
	//		if sum > 1000000000000 {
	//
	//			fmt.Println("end", i)
	//			break
	//		}
	//		fmt.Println("counting")
	//	}
	//}()
	//time.Sleep(100e9)
	//	go func() {
	//		ch <- 1
	//		for i := 0; i < 10000; i++ {
	//			sum += i
	//		}
	//		time.Sleep(1e9)
	//	}()
	//	<-ch
	forumChannel()
}
