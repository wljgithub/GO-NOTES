package main

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func toZip(w io.Writer) {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme", "hello world in go"},
		{"main.go", "package main\nfunc main(){fmt.Println()}"}}
	for _, file := range files {
		f, err := writer.Create(file.Name)
		check(err)
		_, err = f.Write([]byte(file.Body))
		check(err)
	}
	err := writer.Close()
	check(err)

	buf.WriteTo(w)
}
func uploadZipFolder(url string) {
	// 创建表单文件
	// CreateFormFile 用来创建表单，第一个参数是字段名，第二个参数是文件名
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	_, err := writer.CreateFormFile("uploadfile", "file.zip")
	if err != nil {
		log.Fatalf("Create form file failed: %s\n", err)
	}

	toZip(buf)
	// 从文件读取数据，写入表单
	//srcFile, err := os.Open("/Users/xinda/Desktop/GOPATH/src/infoStatistic/file.zip")
	//if err != nil {
	//	log.Fatalf("%Open source file failed: s\n", err)
	//}
	//defer srcFile.Close()
	//_, err = io.Copy(formFile, srcFile)
	//if err != nil {
	//	log.Fatalf("Write to form file falied: %s\n", err)
	//}

	// 发送表单
	contentType := writer.FormDataContentType()
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	_, err = http.Post(url, contentType, buf)
	if err != nil {
		log.Fatalf("Post failed: %s\n", err)
	}

}

func main() {
	url := "http://localhost:9090/upload"
	uploadZipFolder(url)

}
