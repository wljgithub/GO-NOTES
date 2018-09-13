package main

import (
	"reflect"
	"fmt"
)

type People struct{
	Name string
	Age int
}

func convStructToSlice(i interface{})(slice []interface{})  {

		v := reflect.ValueOf(i)
		count := v.NumField()
		for i := 0; i < count; i++ {
			f := v.Field(i)
			switch f.Kind() {
			case reflect.String:
				fmt.Println(f.String())
				slice= append(slice,v.Field(i).String())
			case reflect.Int:
				fmt.Println(f.Int())
				slice= append(slice,v.Field(i).Int())
			}
		}
	return
}
func main()  {
	//var m =make( map[interface{}]interface{})
	//m["1"]=1
	p:=People{"jack",23}
	//t:=reflect.TypeOf(p)
	//fmt.Println(t.Field(0).Name)
	//fmt.Println(t.Field(0).Index)

	//switch t.Kind(){
	//case reflect.String:
	//	m[t.Field(0).Tag]=t.Field(0).Name
	//case reflect.Int:
	//	m[t.Field(1).Tag] = t.Field(1).Name
	//}
	slice:=convStructToSlice(p)
	fmt.Println(slice)
}
