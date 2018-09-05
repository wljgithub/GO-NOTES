package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string
	Age  int
}

func iterOverStruct() {

	v := reflect.ValueOf(People{"jack", 23})
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

func convStructToSlice(s interface{}) (sli []interface{}) {
	v := reflect.ValueOf(s)
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Int:
			sli = append(sli, f.Int())
		case reflect.String:
			sli = append(sli, f.String())

		}
	}
	return
}

func main() {

	m := People{"jack", 23}
	s := convStructToSlice(m)
	fmt.Println(s...)

}
