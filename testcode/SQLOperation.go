package main

import (
		_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"reflect"
	"fmt"
	"log"
)

var  (
)

func check(err error)  {
	if err!=nil {
		panic(err)
	}
}

func convStructToSlice(i interface{})(slice []interface{})  {
	v := reflect.ValueOf(i)
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println(f.String())
			slice= append(slice,v.Field(i).String())
		case reflect.Int:
			fmt.Println(f.Int())
			slice= append(slice,v.Field(i).Int())
		case reflect.Bool:
			slice = append(slice,v.Field(i).Bool())
		case reflect.Float64:
			slice = append(slice,v.Field(i).Float())
		}
	}
	return
}

type MyDB struct {
	db *sql.DB
}

func (this *MyDB)init()  {
	db,err := sql.Open("mysql","root:password@/testDB?charset=utf8")
	if err!=nil {
		panic(err)
	}
	this.db = db
}
func (this *MyDB)insert(table string,values ...interface{}) (error) {
	tx,err:=this.db.Begin()
	check(err)

	query:=fmt.Sprintf("insert into %s values(?,?,?,?,?);",table)
	_,err=tx.Exec(query,values...)
	if err!=nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (this *MyDB)query()  {
	rs,err:=this.db.Query("select * from student_infos;")
	if err!=nil {
		log.Fatal(err)
	}
	type student struct {
		studentID int
		name      string
		age       int
		sex			int
		insertTime string
	}
	someOne := student{}
	for rs.Next(){
		err:=rs.Scan(&someOne.studentID,&someOne.name,&someOne.age,&someOne.sex,&someOne.insertTime)
		fmt.Println(someOne,err)
	}
}
func main() {

	//mydb:=new(MyDB)
	//mydb.init()  // 不同的环境，记得修改init() 下的数据库url
	//defer mydb.db.Close()
	//
	//aStudent := struct {
	//	student_id string
	//	name string
	//	age string
	//	sex string
	//	time string
	//}{
	//	student_id:"666621",
	//	name :"jack",
	//	age :"23",
	//	sex:"0",
	//	time:time.Now().Format("2006-01-02 15:04:05"),
	//}
	//slice:=convStructToSlice(aStudent)
	//fmt.Println("slice:",slice)
	//slice = []interface{}{555554,"jack",23,true,"2006-01-02 15:04:05"}
	//err:=mydb.insert("student_infos",slice...)

	//check(err)
	//mydb.query()

	//var(
	//	name string
	//	age int
	//)
	//db.QueryRow("select name,age from student_infos").Scan(&name,&age)
	//
	//fmt.Println(name,age)

	db,err := sql.Open("mysql","root:password@/testDB?charset=utf8")
	defer db.Close()
	if err!=nil {
		panic(err)
	}
	rs,err:=db.Exec("INSERT INTO testDB.student_infos (student_id,sex) values(123321,true)")
	fmt.Println(rs,err)

	//s1:=[]interface{}{"32144","111","11","11","2006-01-02 15:04:05",}
	//s2:=[]interface{}{"2223243","222","22","22","2006-01-02 15:04:05",}
	//
	//s1=append(s1,s2...)
	//s1 = nil
	//_,err =db.Exec("insert into student_infos values(?,?,?,?,?) , (?,?,?,?,?)",s1...)
	//if err!=nil {
	//	panic(err)
	//}1536831566

	//queryStr :=`select  a.platform,b.value4  from emoa_statisc.v4_preset_client  as a,emoa_statisc.v4_record_client as b where a.env_id = b.env_id and b.value4 = 1 and b.ctimex between 1536831560 and 1536831570;
	//`
	//rs,err:=db.Query(queryStr)
	//if err!=nil {
	//	panic(err)
	//}
	//var platform,msgType string
	//for rs.Next() {
	//	rs.Scan(&platform,&msgType)
	//	fmt.Println(platform,msgType)
	//}

	//fmt.Println(strings.Repeat("(?,?,?,?,?),",5))
	//rs,err:=db.Exec("show create table v4_preset_client;")
	//fmt.Println(rs)
	//fmt.Println(err)
	var g []int
	g=[]int{1,2}
	delete(g,1)
	fmt.Println()
}
