package main

import "fmt"

type Interge int

// 修改对象的值，必须用指针，如果改为func (a Interger)Add..则a的值还是为1
func (a *Interge) Add(b Interge) {
	*a += b
}
func main() {
	var a Interge = 1
	var b Interge = 2
	a.Add(b)
	fmt.Println(a)
}
