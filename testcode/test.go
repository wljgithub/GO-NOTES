package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
		"reflect"
)


func md5Encrypt(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
func DecodeHex(str string) string {
	return hex.EncodeToString([]byte(str))
}

func Q1()  {
	const (
		a = iota
		b
		c
		d
	)
	var v =[...]string{
		a :"a",
		b:"b",
		c:"c",
		d:"d",
	}
	fmt.Println(v==[4]string{"a","b","c","d"})
}
func Q2()  {
	//var i interface{} ="z"[0]
	// range遍历字符串时，返回的是rune,rune是int32的别名
	// rune也可以用单引号缩写 'a'

	for _,v:=range "1"{
		fmt.Println(v == '1')
		fmt.Println('哈'-'嘻')
		fmt.Println()
	}

}
func testFalltrough()  {
	s:=2
	switch s {
	case 2:
		fmt.Println("is 2")
	fallthrough

	case 4:
		fmt.Println("2 or 3")

	}
}

func iterOverStruct()  {
	type People struct {
		Name string
		Age int
	}
	v := reflect.ValueOf(People{"jack",23})
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println(f.String())
		case reflect.Int:
			fmt.Println(f.Int())
		}
	}


}
func main() {
	//sevenDayLater:=time.Now().AddDate(0,0,7).Format("2006-01-02.txt")
	//fmt.Println(sevenDayLater)
	//Q2()
	//testFalltrough()
	//iterOverStruct()
	var a uint32 = 1
	var b uint32	= 2
	fmt.Println(a&b !=0)
}
