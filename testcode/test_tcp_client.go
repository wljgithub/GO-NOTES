package main

import (
	"net"
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func tcpClient()  {
	addr, err := net.ResolveTCPAddr("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	log.Printf("len:%v,content:%s", n, buf)
}
func main() {
	for i:=0;i<300 ;i++  {
		go tcpClient()
		fmt.Println(i)
	}
	fmt.Println()
}
