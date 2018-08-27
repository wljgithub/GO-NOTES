package main

import (
	"bytes"
	"compress/gzip"
	"os"
	"log"
	"fmt"
	"io/ioutil"
)

func check(err error)  {
	if err!=nil {
		log.Fatal(err)
	}
}
func generateGzip(fileName string)  {
	buf:=new(bytes.Buffer)
	writer:=gzip.NewWriter(buf)
	writer.Write([]byte("duo la a meng !"))
	defer writer.Close()

	file,err:=os.OpenFile(fileName,os.O_RDONLY|os.O_CREATE,0777)
	defer file.Close()
	check(err)
	buf.WriteTo(file)
}

func readGzipFile(fileName string)  {
	file,err:=os.OpenFile(fileName,os.O_RDONLY,0777)
	check(err)

	reader,err:=gzip.NewReader(file)
	check(err)
	fmt.Println(string(reader.Extra))

}


func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(in)
	if err != nil {
		writer.Close()
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}
func main() {
	fileName:="./1.gzip"
	//generateGzip(fileName)
	readGzipFile(fileName)

}
