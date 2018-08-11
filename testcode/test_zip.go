package main

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
)

func main() {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)

	// 创建一个压缩文档
	w := zip.NewWriter(buf)

	//rid	record id,该记录的唯一标识并固定不变。
	//key	统计类型	cLogin
	//time	触发时间,毫秒	1478156937000
	//devId	设备Id	dasfsaf
	//buin	企业号	12345678
	//gid	用户id	130012
	//userName	用户账号	cs1
	//os	操作系统	Android
	//osVer	操作系统版本号	6.0
	//plat	客户端类型	2
	//cVer	客户端版本号	2.5.63
	//dict	统计字典	spend%3D90%26result%3D0
	// 将文件加入压缩文档

	var files = []struct{ Name, Body string }{{"readme.txt", "rid=123&key=cLogin&time=1478156937000&devId=dasfsaf&buin=12345678&gid=100037&userName=cs1&os=Android&osVer=6.0&plat=2&cVer=2.5.63&dict=spend%3D90%26result%3D0"}}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println(string(buf.Bytes()))
	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 将压缩文档内容写入文件
	f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	buf.WriteTo(f)
}
