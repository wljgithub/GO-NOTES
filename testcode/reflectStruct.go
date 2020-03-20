package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
}
type Bar struct {
	Name string
	Age int
}

//用于保存实例化的结构体对象
var regStruct map[string]interface{}

func main() {
	str := "Bar"
	if regStruct[str] != nil {
		t := reflect.ValueOf(regStruct[str]).Type()
		fmt.Println(t)
		v := reflect.New(t).Elem()
		fmt.Println(v)
	}

}

func init() {
	regStruct = make(map[string]interface{})
	regStruct["Foo"] = Foo{}
	regStruct["Bar"] = Bar{"jack",23}
}