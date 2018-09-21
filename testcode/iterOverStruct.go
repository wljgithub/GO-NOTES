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
func diveStruct(a,b interface{})interface{}{
	v1 := reflect.ValueOf(a).Elem()
	v2 := reflect.ValueOf(b).Elem()
	count := v1.NumField()
	for i:=0;i<count;i++{
		v1.Field(i).Set(reflect.ValueOf(v1.Field(i).Float()/v2.Field(i).Float()))
		fmt.Println( v1.Field(i).Float()/v2.Field(i).Float())
	}
	fmt.Println(a,v1)
	return v1
}
func main()  {
	//var m =make( map[interface{}]interface{})
	//m["1"]=1
	//p:=People{"jack",23}
	//
	//slice:=convStructToSlice(p)
	//fmt.Println(slice)
	//
	var (
	//	a = struct {
	//		num1 float64
	//		num2 float64
	//	}{10,20}
		b = struct {
			num1 float64
			num2 float64
		}{5,5}
	)
	//diveStruct(a,b)

	s := reflect.ValueOf(&b).Elem()  // 反射获取测试对象对应的struct枚举类型
	//
	//s.Field(0).SetFloat(123) // 内置常用类型的设值方法，利用Field序号get

	fmt.Println(b)
	//sliceValue := reflect.ValueOf([]int{1, 2, 3}) // 这里将slice转成reflect.Value类型; 用于下一步赋值
	s.FieldByName("num1").Set(reflect.ValueOf(1.1))  // 使用
	fmt.Println(b)


}
