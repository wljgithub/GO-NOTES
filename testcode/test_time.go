package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	p:=fmt.Println

	// 当前时间
	p(time.Now().Format("2006-01-02 15:04:05"))

	p(time.Now().Sub(time.Now()).Hours())
}
