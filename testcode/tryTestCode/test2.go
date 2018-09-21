package main

import (
	"fmt"
	"time"
)

func main()  {
	var m = map[string]string{"a":"1"}
	var m1 = m

	m1["2"]="b"
	fmt.Println(m,m1)

	today,err:=time.Parse("2006-01-02 ",time.Now().Format("2006-01-02 "))
	if err!=nil {
		panic(err)
	}

	fmt.Println(today.Unix(),today.AddDate(0,0,1).Unix())
	fmt.Println(today.Format("2006-01-02 15:04:05"),today.AddDate(0,0,1).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Unix(),1537413577,1537401600 ,1537488000)
}
