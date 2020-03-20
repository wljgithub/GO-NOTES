package main

import (
	"net"
	"fmt"
	)

func main() {
	inte,err:=net.Interfaces()
	addr,err:=inte[0].Addrs()
	fmt.Println(addr,err,inte)
}
