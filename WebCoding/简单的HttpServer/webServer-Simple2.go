// 本来想从url中提取路径，然后输出的，设置路由的时候不知道怎么匹配路径，不纠结这个了

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

// 写一个函数，处理HTTP-get请求
func autoSayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器的打印信息
	fmt.Println("path", r.URL.Path)

	// 从URL提取 NAME
	//	Name := r.URL.Path

	// 在 NAME前加 hello，返回给客户端
	fmt.Fprintf(w, "Hello，Jack") //这些写入到w的是输出给客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/jack", autoSayhelloName)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
