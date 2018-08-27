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
	"net"
	"time"
	"log"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		log.Println("get a connect")
		daytime := time.Now().String()
		log.Println(daytime)
		conn.Write([]byte(daytime),) // don't care about return value
		log.Println("write msg success!")
		//conn.Close()                // we're finished with this client
	}
}
func checkError(err error) {
	if err!=nil {
		log.Fatal(err)
	}
}