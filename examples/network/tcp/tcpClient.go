//func check(err error)  {
//	if err!=nil {
//		fmt.Fprintf(os.Stderr,"%s",err.Error())
//	}
//}
//func main() {
//	addr,err:=net.ResolveTCPAddr("tcp4","127.0.0.1:8090")
//	check(err)
//
//	localAddr,err:=net.ResolveTCPAddr("tcp4","127.0.0.1:8080")
//	check(err)
//	conn,err:=net.DialTCP("tcp4",localAddr,addr)
//
//
//	//i,err:=conn.ReadFrom(os.Stdin)
//	//fmt.Printf("read %d byte from stdin",i)
//	conn.Write([]byte("jack"))
//
//	result,err:=ioutil.ReadAll(conn)
//	check(err)
//
//	fmt.Println(string(result))
//}

package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:6565")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, addr)
	defer conn.Close()
	checkError(err)

	input:=bufio.NewScanner(os.Stdin)
	for {
		input.Scan()
		//input, err := ioutil.ReadAll(os.Stdin)
		//checkError(err)

		_,err=conn.Write(input.Bytes())
		checkError(err)

		bytes:=make([]byte,20)
		_,err:=conn.Read(bytes)

		checkError(err)
		fmt.Println(string(bytes))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
