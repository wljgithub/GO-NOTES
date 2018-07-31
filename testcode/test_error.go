package main

import (
	"github.com/kataras/iris/core/errors"
	"fmt"
)

func IsInt(par interface{})(b bool,err error  ){
	switch par.(type) {
	case int:
		b,err= true,nil
	default:
		b,err=false,errors.New("the value is not int type")
	}
	return
}
func main()  {
	b,err:=IsInt(0)
	if err!=nil{
		panic(err)
	}else{
		fmt.Println(b)
	}
}