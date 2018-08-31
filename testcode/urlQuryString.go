package main

import (
	"fmt"
	"net/url"
)

func main()  {
	fmt.Println(url.QueryUnescape("name%3D%E5%B7%AB%E5%87%8C%E5%81%A5"))
	userInfo:=url.UserPassword("jack","*****")
	fmt.Println(userInfo.Username())
	fmt.Println(userInfo.Password())


}