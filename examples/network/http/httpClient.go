package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"fmt"
	"bytes"
	"mime/multipart"
	"io"
	"compress/gzip"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
func httpCustoGet(url string) {
	req, err := http.NewRequest("GET", url, nil)
	checkerr(err)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.84 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	checkerr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkerr(err)
	log.Println(string(body))

}

func httpCustoPost(url string) {

	// 构造请求
	req, err := http.NewRequest("POST", url,nil)
	checkerr(err)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.84 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	checkerr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkerr(err)

	log.Printf("%s", body)
}

func postFormValue() {
	resp, err := http.PostForm("http://www.httpbin.org/post", url.Values{"name": {"jack", "huimin"}, "age": {"23", "21"}})
	checkerr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkerr(err)
	log.Printf("%s", body)
// result :
	//"form": {
	//	"age": [
	//	"23",
	//		"21"
	//	],
	//	"name": [
	//	"jack",
	//		"huimin"
	//	]
	//},
}
func httpGet() {
	res, err := http.Get("http://www.httpbin.org/get")
	checkerr(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkerr(err)

	log.Printf("%s", body)
}
func toGzip(w io.Writer) {

	buf := new(bytes.Buffer)
	writer := gzip.NewWriter(buf)


	n,err:=writer.Write([]byte("duo la a meng !\n"))
	log.Println("data len:",n)
	if err!=nil {
		log.Fatal(err)
	}
	err = writer.Close()
	if err!=nil {
		log.Fatal(err)
	}
	//_, err = buf.WriteTo(w)
	buf.WriteTo(w)


}
func uploadFile(url string) {

	// CreateFormFile 用来创建表单，第一个参数是字段名，第二个参数是文件名
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	_, err := writer.CreateFormFile("file", "file.gzip")
	if err != nil {
		log.Fatalf("Create form file failed: %s\n", err)
	}

	toGzip(buf)
	// 从文件读取数据，写入表单

	// 发送表单
	contentType := writer.FormDataContentType()
	log.Println(contentType)
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	resp, err := http.Post(url, contentType, buf)
	if err != nil {
		log.Fatalf("Post failed: %s\n", err)
	}
	data,err:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(data),err)
}


func main() {
	// 普通的get请求
	httpGet()

	//// 自定义user-agent的 get请求
	httpCustoGet("http://www.httpbin.org/get")

	// 提交表单数据的post请求，但不知道提交的表单是什么格式的数据
	postFormValue()

	// 自定义user-agent的 post请求
	url := "http://www.httpbin.org/post"
	httpCustoPost(url)


	// 文件上传
	uploadFile(url)

}
