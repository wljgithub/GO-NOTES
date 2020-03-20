package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3307)/blog?charset=utf8")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	var platform, msgType, msgCount, userCount ,receipt int
	//各平台的各类型消息使用次数和人数
	// 这里好像有点问题，总次数不需要去重的，但是这里distinct(r.gid)根据gid去重了
	rs, err := db.Query(`select platform,value4,count(value9=1 ),count(0),count(distinct(r.gid)) from v4_record_client r left  join v4_preset_client p on r.env_id = p.env_id where r.rkey=?  group by platform,value4`,
		"cSent")
	if err != nil {
		panic(err)
	}
	for rs.Next() {
		err:=rs.Scan(&platform, &msgType,&receipt, &msgCount, &userCount)
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println(platform, msgType,receipt, msgCount, userCount)
	}
	rs.Close()
	//`delete  from v4_record_client r  v4_preset_client p wherer r.env_id = p.env_id and gid = 12345`
	//_,err =db.Exec(	`DELETE v4_preset_client ,v4_record_client from v4_preset_client left JOIN v4_record_client ON v4_preset_client .env_id=v4_record_client.env_id WHERE v4_record_client.gid=22334;`)
	//if err!=nil {
	//	panic(err)
	//}
	//`DELETE v4_record_client,v4_preset_client from v4_record_client LEFT JOIN v4_preset_client ON v4_record_client.gid=v4_preset_client.gid WHERE v4_preset_client.id=25`
//
}
