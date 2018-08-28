package main

import (
	"errors"
)

func Division(x,y float64)(float64,error)  {
	if y==0 {
		return 0,errors.New("被除数不能为零")
	}
	return x/y,nil
}
