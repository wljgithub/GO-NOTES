package main
import (
	"fmt"
	"io"
	"net"
	"os"
	"log"
)

func check(err error)  {
	if err!=nil {
		log.Fatal(err)
	}
}
func main() {
	addr,err:=net.ResolveTCPAddr("tcp",":3333")
	check(err)

	l,err:=net.ListenTCP("tcp",addr)
	check(err)

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		go handConn(conn)
		//logs an incoming message


	}
}

func handConn(conn net.Conn)  {
	fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
	// Handle connections in a new goroutine.
	//go handleRequest(conn)
	raw :=make([]byte,20)
	n,err:=conn.Read(raw)
	log.Println("len:",n,"receive:",string(raw))
	_,err=conn.Write([]byte(fmt.Sprintln("hello",string(raw))))
	if err!=nil {
		log.Fatal(err)
	}
	conn.Close()

}
func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}