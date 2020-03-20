package main

import (
	"net/http"
	"log"
	)

func sayHello(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	w.Write([]byte("hello "+string(r.URL.Path)))
	//w.Write([]byte("hello "+string(r.URL.Scheme)))
}

func main() {
	// "/" 可以匹配任何路径
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal(err)
	}
	
}
