/*
go程序是通过package来组织的，除了main包一位，其他的包都会生成*.a文件（也就是包文件）放在$GOPATH/pkg/$GOOS_$GOARCH中

*/
package main

//为了打印helloworld，调用了Print，这个函数来自fmt包
import "fmt"

//
func main() {
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
}
