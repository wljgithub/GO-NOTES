package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}
func (p *Page)loadPage()(*Page,error){
	var title = p.Title + ".txt"
	var body,err = ioutil.ReadFile(title)
	return &Page{Title:title,Body:body},err
}
func (p *Page)save()error{
	return ioutil.WriteFile(p.Title,p.Body,0600)
}

func viewHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hi,are you call %s",r.URL.Path[1:])
}
func main() {
		http.HandleFunc("/view/",viewHandler)
		http.ListenAndServe(":8080",nil)
}
