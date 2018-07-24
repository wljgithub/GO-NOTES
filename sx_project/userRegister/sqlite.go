package main

import (
	//	"database/sql"
	//	"fmt"
	//	"time"

	_ "github.com/mattn/go-sqlite3"
)

//var (
//	rejectUserTable = `CREATE TABLE userInfo (

//	        account    STRING (20)         PRIMARY KEY     ASC ON CONFLICT ABORT    NOT NULL    UNIQUE,

//	        name        STRING (32)         UNIQUE    NOT NULL,

//	        sex            INTEGER (2)         UNIQUE,

//	        email        STRING (24)         UNIQUE    NOT NULL,

//	        password  STRING(24)         NOT NULL,

//	        cellNum    STRING (11),

//	        teleNum   STRING (12),

//	        status        INTEGER (2)        UNIQUE,

//	        deptId       STRING (4),

//	        date          DATE (10)            UNIQUE

//	);`
//	userInfo = `CREATE TABLE rejectUserTable (

//	    account        STRING (20)  PRIMARY KEY    NOT NULL    UNIQUE,

//	    date              DATE (10)     NOT NULL,

//	    reson            STRING (280)

//	);`
//)
//var db, err = sql.Open("sqlite3", "test.db")

//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//func createTable(s string) {
//	_, err = db.Exec(s)
//	fmt.Println(err)
//}
//func main() {

//	//	//创建两张表
//	createTable(rejectUserTable)
//	createTable(userInfo)

//	//	db.Exec("insert into test(id,age) values(?,?)", "first", 18)

//	//插入数据
//	stmt, err := db.Prepare("INSERT INTO rejectUserTable(account, date, reson) values(?,?,?)")
//	checkErr(err)

//	num := []interface{}{"李十四", "2018-07-24", "测试插入数据"}
//	res, err := stmt.Exec(num...)
//	checkErr(err)

//	id, err := res.LastInsertId()
//	checkErr(err)

//	fmt.Println(id)

//	//	stmt, err = db.Prepare("INSERT INTO userInfo values(?,?,?,?,?,?,?,?,?,?)")
//	//	//	stmt, err = db.Prepare("INSERT INTO userInfo (account,name)values(?,?)")
//	//	checkErr(err)

//	//	strs := []interface{}{"jackAccount", "jack", "0", "332150887@qq.com", "password", "15913266105", "123456789101", "01", "aprt", "2018-7-21"}
//	//	//	strs := []interface{}{"jackwu", "jack"}
//	//	res, err = stmt.Exec(strs...)
//	//	checkErr(err)
//	//	id, err = res.LastInsertId()
//	//	checkErr(err)

//	//	fmt.Println(id)

//	//	//更新数据
//	//	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
//	//	checkErr(err)

//	//	res, err = stmt.Exec("astaxieupdate", id)
//	//	checkErr(err)

//	//	affect, err := res.RowsAffected()
//	//	checkErr(err)

//	//	fmt.Println(affect)

//	//	//查询数据
//	rows, err := db.Query("SELECT * FROM rejectUserTable")
//	//	checkErr(err)

//	//	for rows.Next() {
//	//		var uid int
//	//		var username string
//	//		var department string
//	//		var created time.Time
//	//		err = rows.Scan(&uid, &username, &department, &created)
//	//		checkErr(err)
//	//		fmt.Println(uid)
//	//		fmt.Println(username)
//	//		fmt.Println(department)
//	//		fmt.Println(created)
//	//	}
//	for rows.Next() {
//		var account string
//		var date string
//		var reson string
//		err = rows.Scan(&account, &date, &reson)
//		checkErr(err)
//		fmt.Println(account)
//		fmt.Println(date)
//		fmt.Println(reson)
//	}

//	//	//删除数据
//	//	stmt, err = db.Prepare("delete from userinfo where uid=?")
//	//	checkErr(err)

//	//	res, err = stmt.Exec(id)
//	//	checkErr(err)

//	//	affect, err = res.RowsAffected()
//	//	checkErr(err)

//	//	fmt.Println(affect)

//	defer db.Close()

//}
