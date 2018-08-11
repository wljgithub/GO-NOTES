package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func uploadZipFolder(url string) {
	// 创建表单文件
	// CreateFormFile 用来创建表单，第一个参数是字段名，第二个参数是文件名
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile("uploadfile", "file.zip")
	if err != nil {
		log.Fatalf("Create form file failed: %s\n", err)
	}

	// 从文件读取数据，写入表单
	srcFile, err := os.Open("/Users/xinda/Desktop/GOPATH/src/infoStatistic/file.zip")
	if err != nil {
		log.Fatalf("%Open source file failed: s\n", err)
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Fatalf("Write to form file falied: %s\n", err)
	}
	//_,err=formFile.Write([]byte("asdasdasdasd"))

	// 发送表单
	contentType := writer.FormDataContentType()
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	_, err = http.Post(url, contentType, buf)
	if err != nil {
		log.Fatalf("Post failed: %s\n", err)
	}

}

func main() {
	url := "http://localhost:8080/upload"
	uploadZipFolder(url)

}
