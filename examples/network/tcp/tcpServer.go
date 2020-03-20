//package main
//
//import (
//	"fmt"
//	"net"
//	"os"
//)
//
//func check(err error) {
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "%s", err.Error())
//	}
//}
//
//func main() {
//	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
//	check(err)
//
//	listener, err := net.ListenTCP("tcp4", addr)
//	check(err)
//	fmt.Println("tcp server start listening!")
//	conn, err := listener.Accept()
//	check(err)
//
//	//var b []byte
//	//conn.Read(b)
//
//	returnData := fmt.Sprintf("hello")
//	conn.Write([]byte(returnData))
//	conn.Close()
//
//}
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:6565")
	checkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)
	log.Println("tcp server is running...")
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("connect error", err.Error())
			break
		}
		log.Println("connect:", conn.RemoteAddr(), "-->", conn.LocalAddr())
		go handConnect(conn)
	}
}

func handConnect(conn net.Conn) {
	for {

		buf := make([]byte, 20)
		_, err := conn.Read(buf)
		checkError(err)
		returnData := []byte(fmt.Sprintf("hello,%s!", buf))
		_,err=conn.Write(returnData)
		if err!=nil {
			fmt.Println("occur err:",err.Error())
			break
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
