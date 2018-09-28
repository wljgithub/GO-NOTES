package main
import (
		"fmt"
	"math/rand"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strings"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var HashVal []string
var Num =1000
var rkeys []string
var eventType []int
func GetRandomString(length int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}

	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


func generaPlatform(num int,platform *[]int){
	for i:=0;i<num ;i++  {
		if i<num/4{
			*platform = append(*platform,1)
		}
		if i>=(num/4)&&i<(num*2/4){
			*platform = append(*platform,2)
		}
		if i>=num*2/4&&i<num*3/4{
			*platform = append(*platform,3)
		}
		if i>=num*3/4{
			*platform = append(*platform,4)
		}

	}
	//fmt.Println(*platform)
}
func generateRkey(num int ){
	//for i:=0;i<num/4 ;i++  {
	//	rkeys= append(rkeys,"cLogin")
	//}
	for i:=0;i<num ;i++  {
		rkeys= append(rkeys,"cSent")
	}

}
func generateEventType(num int)  {
	for i:=0;i<30;i++{
		eventType = append(eventType,1)
		eventType = append(eventType,3)
		eventType = append(eventType,4)
		eventType = append(eventType,5)

		eventType = append(eventType,7)
		eventType = append(eventType,12)
		eventType = append(eventType,13)

	}
	for i:=30;i<num ;i++  {
		eventType = append(eventType,0)
	}
}
func main()  {
	db,err := sql.Open("mysql","root:password@/emoa_statisc?charset=utf8")
	defer db.Close()
	if err!=nil {
		panic(err)
	}
	// 生成 xxx个随机字符串
	for i:=0;i<Num ;i++  {
		str:=GetRandomString(len("09b56806c44069177b9c8f922fc12bsq"))
		HashVal = append(HashVal,str)
	}

	// 生成查询的sql语句
	str:=strings.Repeat("?,",23)
	query:=fmt.Sprintf("INSERT INTO v4_preset_client values(%s?)",str)

	platform:=[]int{}
	generaPlatform(Num,&platform) // 伪造500个移动端，250个pc端登录

	gid:=[]int{12345,54321,22334}
	var preset = []interface{}{"09b56806c44069177b9c8f922fc12bsq", "72907129", "331532", "866334031878826", "", "", "", "1", "3.2.71", "1", "NRD90M", "Xiaomi", "mido", "Android", "24", "64", "1920", "1080", "46000", "0", "armeabi-v7a", "", "", "0",
	}

	for i:=0;i<Num;i++  {

		preset[0]= HashVal[i]// 随机字符串 --envID
		preset[2] =gid[r.Intn(len(gid))] // gid  随机数字
		preset[7] = platform[i] // platform 1 2 3 4

		_,err:=db.Exec(query,preset...)
		time.Sleep(time.Millisecond)
		if err!=nil {
			fmt.Println(err)
		}
	}

	str =strings.Repeat("?,",22)
	query =fmt.Sprintf("INSERT INTO v4_record_client values(%s?)",str)
	event := []interface{}{"2", "09b56806c44069177b9c8f922fc8fafa", "331532", "72907129", "cLogin", "1536831570", "2018-09-13 15:58:12", "0", "0", "0", "0", "4", "0", "0", "0", "0", "0", "0", "", "", "", "", ""}

	//rkey:=[]string{"cSent","cLogin"}
	generateRkey(Num)
	generateEventType(Num)
	for i:=0;i<Num ;i++  {
		event[0] = i   // rid
		event[1] = HashVal[i]  //envId
		event[2] =gid[r.Intn(len(gid))]  //gid
		event[4] = rkeys[i] //rkey

		event[5] =time.Now().UnixNano()  //ctimex
		event[11]   = eventType[i]   // 消息类型
		event[16]  = false  // 回执

		_,err:=db.Exec(query,event...)
		time.Sleep(time.Millisecond)
		if err!=nil {
			fmt.Println(err)
		}
	}
}
