package main

import (
		"io"
	"fmt"
	"crypto/md5"
		"crypto/sha1"
	"crypto/sha256"
)

func md5EncryptAddSult(passwd string) string {
	h := md5.New()
	io.WriteString(h, passwd)

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 :=fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf("%x",h.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last :=fmt.Sprintf("%x", h.Sum(nil))

	return last
}

func md5Encrypt(passwd string)string  {

	//import "crypto/md5"
	h := md5.New()
	io.WriteString(h, passwd)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func sha1Enctypt(passwd string)string  {

	h := sha1.New()
	io.WriteString(h, passwd)
	return fmt.Sprintf("% x", h.Sum(nil))
}

func sha256Encrypt(passwd string)string  {

	h:=sha256.New()
	io.WriteString(h,passwd)
	return fmt.Sprintf("%x",h.Sum(nil))
}
func main()  {

	p:=fmt.Println
	passwd := "password"
	//md5加盐
	p(md5EncryptAddSult(passwd))

	//md5
	p(md5Encrypt(passwd))

	//sha1
	p(sha1Enctypt(passwd))

	p(sha256Encrypt(passwd))

}
