package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime"
)

func md5Encrypt(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
func DecodeHex(str string) string {
	return hex.EncodeToString([]byte(str))
}

func Q1() {
	const (
		a = iota
		b
		c
		d
	)
	var v = [...]string{
		a: "a",
		b: "b",
		c: "c",
		d: "d",
	}
	fmt.Println(v == [4]string{"a", "b", "c", "d"})
}
func Q2() {
	//var i interface{} ="z"[0]
	// range遍历字符串时，返回的是rune,rune是int32的别名
	// rune也可以用单引号缩写 'a'

	for _, v := range "1" {
		fmt.Println(v == '1')
		fmt.Println('哈' - '嘻')
		fmt.Println()
	}

}
func testFalltrough() {
	s := 2
	switch s {
	case 2:
		fmt.Println("is 2")
		fallthrough

	case 4:
		fmt.Println("2 or 3")

	}
}
func printOut(n int) {
	if n >= 10 {
		printOut(n / 10)
	}
	fmt.Println(n % 10)
}

//func test(i interface{}) {
//	fmt.Println(i)
//	for _, v := range i {
//		fmt.Println(v)
//	}
//}

func main() {
	//sevenDayLater:=time.Now().AddDate(0,0,7).Format("2006-01-02.txt")
	//fmt.Println(sevenDayLater)
	//Q2()
	//testFalltrough()
	//iterOverStruct()
	var a uint32 = 1
	var b uint32 = 2
	fmt.Println(a&b != 0)
	fmt.Printf("%q", 'n')
	runtime.GOMAXPROCS(2)
	fmt.Println(runtime.NumCPU())

	fmt.Println(1234 % 10)
	printOut(1234)

	type MyStruct struct {
		Name string
	}
	type Other struct {
		Age int
	}
	var i interface{} = MyStruct{"jack"}
	v, ok := i.(Other)
	fmt.Println(v, ok)

	var slice = []int{1, 2, 3}
	fmt.Println(slice)
	slice = []int{}
	fmt.Println(slice)

	fmt.Println(10<<20 == 10*1024*1024)

	var err error
	fmt.Println(err == nil)

	var m map[string]string
	fmt.Println(m == nil)

	var s []string
	fmt.Println(s == nil)

	var okk bool
	fmt.Println(okk)


	nn:=[]int{1,2,3}
	mm:=[]int{4,5,6}
	nn = append(nn,mm...)
	fmt.Println(nn)

	mmm:=make([]interface{},10)
	fmt.Println(mmm[1])
}
