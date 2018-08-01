package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
	"log"
)

func simpleGet()  {
	url := "http://httpbin.org/get"
	fmt.Println("URL:",url)

	resp,err:=http.Get(url)
	if err!=nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}

func simplePost()  {

		url := "http://httpbin.org/post"
		fmt.Printf("url: %s\n", url)

		res, err := http.Post(
			url,
			"application/x-www-form-urlencoded",
			strings.NewReader("name=jack"))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", body)

}
func customPost()  {
	url := "http://httpbin.org/post"
	fmt.Println("URL:>", url)

	// 字符串拼接
	//	encryMsg := "passwd"
	//	end := fmt.Sprintf(`"encrypt": "%s"}`, encryMsg)

	// 原来 post 的 json 可以这么写，神奇
	jsonStr := []byte(`{"buin":"14998910","appID":"yd2FA9310DC3AA442ABBBF343F891DF3C1"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
func main() {
	//simpleGet()
	simplePost()
	//customPost()
}
