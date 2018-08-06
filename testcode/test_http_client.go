package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
func httpClientGet(url string) {
	req, err := http.NewRequest("GET", url, nil)
	checkerr(err)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.84 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	checkerr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkerr(err)
	log.Panicln(string(body))

}

func httpClientPost(url string) {
	req, err := http.NewRequest("POST", url, nil)
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

}
func httpGet() {
	res, err := http.Get("http://www.httpbin.org/get")
	checkerr(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkerr(err)

	log.Printf("%s", body)
}
func main() {
	// 普通的get请求
	httpGet()

	// 自定义user-agent的 get请求
	httpClientGet("http://www.httpbin.org/get")

	// 提交表单数据的post请求，但不知道提交的表单是什么格式的数据
	postFormValue()

	// 自定义user-agent的 post请求
	httpClientPost("http://www.httpbin.org/post")

	// 差一个post上传文件
	// resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

}
