package main

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"mime/multipart"
	"net/http"
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
		{"zip1", "rid=123&key=cLogin  &time=1478156937000&devId=dasfsaf&buin=12345678&gid=100037&userName=cs1&os=Android&osVer=6.0&plat=2&cVer=2.5.63&dict=spend%3D90%26result%3D0 rid=123&key=cLogin&time=1478156937000&devId=dasfsaf&buin=12345678&gid=100037&userName=cs1&os=Android&osVer=6.0&plat=2&cVer=2.5.63&dict=spend%3D90%26result%3D0\n"},
	}
	for _, file := range files {
		f, err := writer.Create(file.Name)
		check(err)
		for i := 0; i < 2; i++ {

			_, err = f.Write([]byte(file.Body))
			check(err)
		}

	}
	err := writer.Close()
	check(err)

	buf.WriteTo(w)
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
func uploadZipFolder(url string) {
	// 创建表单文件
	// CreateFormFile 用来创建表单，第一个参数是字段名，第二个参数是文件名
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	_, err := writer.CreateFormFile("file", "file.gzip")
	if err != nil {
		log.Fatalf("Create form file failed: %s\n", err)
	}

	//toZip(buf)

	toGzip(buf)
	// 从文件读取数据，写入表单

	// 发送表单
	contentType := writer.FormDataContentType()
	log.Println(contentType)
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	_, err = http.Post(url, contentType, buf)
	if err != nil {
		log.Fatalf("Post failed: %s\n", err)
	}
	//atomic.AddInt32(&count, 1)
	//log.Println(atomic.LoadInt32(&count))
}

func main() {
	url := "http://127.0.0.1:17016/v3/api/jgstatisc/collect.do"
	for i := 0; i < 5; i++ {
		go uploadZipFolder(url)
		time.Sleep(10 * time.Millisecond)

	}

}
