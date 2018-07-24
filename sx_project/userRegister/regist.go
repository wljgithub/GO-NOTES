package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//	"strings"
)

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {
		//从表单提取数据，存入数据库
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		//		`{

		//“errmsg”：""，
		//“code”：1（1为提交成功，2为提交失败）

		//“user”：{
		//	“account”：“1294139563”
		//	“name”：“Tom”
		//    “password”：一次md5 “123456”
		//    “sex”：1（0为男，1为女），
		//	“email”：“1296415646@qq.com”，
		//	“cellNum”：“156269677582”，
		//	“teleNum”：“1319225489”，
		//	“deptId”：“001”,
		//	“registertime”：“2018-07-24  13:25:45”
		//	}

		//}`
	}
}
func adminAudit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method) //获取请求的方法
	// 如果是GET 方法，重定向到管理员审核
	if r.Method == "GET" {

		t, _ := template.ParseFiles("resource/user-table.html")
		log.Println(t.Execute(w, nil))
	} else if r.Method == "POST" {

		//管理员审核页面的HTTP请求，返回json格式的数据
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func adminLogin(w http.ResponseWriter, r *http.Request) {
	//	`
	//input
	//{
	//“name”：“sony”，
	//“password”： 一次md5 “123456”
	//}

	//output
	//{
	//“code”：1，
	//“errmsg”：“用户名或密码错误”
	//}`
	// 如果是get 请求，返回登录页面
	r.ParseForm()
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("resource/login.html")
		log.Println(t.Execute(w, nil))

		//如果是 POST 请求，
	} else if r.Method == "POST" {
		//验证用户输入用户名和密码(放在本地的文件当中),如果成功了，重定向到管理员审核页面
		//		fmt.Println("username:", r.Form["username"])
		//		fmt.Println("password:", r.Form["password"])
		fmt.Println(r.Form)
	}
}
func approve(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		//`input

		//account=abc

		//output
		//{
		//“errmsg”：""，
		//“code”：1（1为提交成功，2为提交失败）
		//“account”："aaa"
		//}`
	}
}
func getAllUser(w http.ResponseWriter, r *http.Request) {

	//从数据库 拉取所有 待审批用户的数据，以json的格式返回
}

func reject(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.Handle("/", http.FileServer(http.Dir("resource/"))) //用户注册页面
	http.HandleFunc("/register", register)                   //接受用户注册的信息，导入数据库

	http.HandleFunc("/login", adminLogin) //get管理员登录页面,post验证管理员登录密码

	http.HandleFunc("/admin/audit", adminAudit) //管理员审核页面 - 用来重定向返回给客户端

	http.HandleFunc("/approve", approve) //审批通过 - 调有度sdk 注册
	http.HandleFunc("/reject", reject)   //审批驳回 - 发post 请求，将用户、日期、原因发到服务器，入库

	http.HandleFunc("/getAllUser", getAllUser) //get 获取所有的用户信息-返回的时候json格式的数据

	//	http.HandleFunc("/getAlldept",getAlldept) //get 获取所有部门信息

	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
