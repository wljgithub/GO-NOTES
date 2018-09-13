package main

import "fmt"

func testRecover(){
	defer func() {
		if r:=recover();r!=nil{
			fmt.Println("err:",r)
		}
	}()
	panic("make err")

}
func main()  {
	testRecover()

}