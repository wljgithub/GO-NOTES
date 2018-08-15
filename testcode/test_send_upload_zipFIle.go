package main

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	count int32
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
		{"zip1", "rid=123&key=cLogin  &time=1478156937000&devId=dasfsaf&buin=12345678&gid=100037&userName=cs1&os=Android&osVer=6.0&plat=2&cVer=2.5.63&dict=spend%3D90%26result%3D0 rid=123&key=cLogin&time=1478156937000&devId=dasfsaf&buin=12345678&gid=100037&userName=cs1&os=Android&osVer=6.0&plat=2&cVer=2.5.63&dict=spend%3D90%26result%3D0"},
		// {"zip2", "rid=123&key=cLogin&time=1478156937000&devId=dasfsaf&buin=12345678&gid=100037&userName=cs1&os=Android&osVer=6.0&plat=2&cVer=2.5.63&dict=spend%3D90%26result%3D0 rid=123&key=cLogin&time=1478156937000&devId=dasfsaf&buin=12345678&gid=100037&userName=cs1&os=Android&osVer=6.0&plat=2&cVer=2.5.63&dict=spend%3D90%26result%3D0"}}
	}
	for _, file := range files {
		f, err := writer.Create(file.Name)
		check(err)
		for i := 0; i < 20000; i++ {

			_, err = f.Write([]byte(file.Body))
			check(err)
		}

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
	atomic.AddInt32(&count, 1)
	log.Println(atomic.LoadInt32(&count))
}

func main() {
	url := "http://127.0.0.1:18090/v4/api/jgstatisc/collect.do"
	for i := 0; i < 2; i++ {
		go uploadZipFolder(url)
		time.Sleep(10 * time.Millisecond)

	}
	select {}
}
