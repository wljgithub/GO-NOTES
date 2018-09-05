package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func generateGzip(fileName string) {
	writer := gzip.NewWriter(nil)
	_, err := writer.Write([]byte("duo la a meng !"))
	check(err)
	defer writer.Close()

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	check(err)
	defer file.Close()

	io.Copy(writer, file)
}

func readGzipFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0777)
	defer file.Close()
	check(err)

	reader, err := gzip.NewReader(file)
	check(err)
	err = reader.Close()
	check(err)

	raw, err := ioutil.ReadAll(reader)
	check(err)
	fmt.Println(string(raw))

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
	fileName := "./1.gzip"
	generateGzip(fileName)
	readGzipFile(fileName)

}
