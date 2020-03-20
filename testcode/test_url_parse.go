package main

import (
	"fmt"
	"net/url"
)

type DataReport struct {
	Rid      string `json:"rid"`
	Key      string `json:"key"`
	Time     string `json:"time"`
	DevId    string `json:"devId"`
	Buin     string `json:"buin"`
	Gid      string `json:"gid"`
	UserName string `json:"userName"`
	Os       string `json:"os"`
	OsVer    string `json:"osVer"`
	Plat     string `json:"plat"`
	CVer     string `json:"cVer"`
	Dict     string `json:"dict"`
}

func queryStrParse(queryStr string) DataReport {
	value, err := url.ParseQuery(queryStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(value.Get("dict"))

	return DataReport{
		Rid:      value.Get("rid"),
		Key:      value.Get("key"),
		Time:     value.Get("time"),
		DevId:    value.Get("devId"),
		Buin:     value.Get("buin"),
		Gid:      value.Get("gid"),
		UserName: value.Get("userName"),
		Os:       value.Get("os"),
		OsVer:    value.Get("osVer"),
		Plat:     value.Get("plat"),
		CVer:     value.Get("cVer"),
		Dict:     value.Get("dict"),
	}
}

func main() {
	queryStr := `rid=123&key=cLogin&time=1478156937000&devId=dasfsaf&buin=12345678&gid=100037&userName=cs1&os=Android&osVer=6.0&plat=2&cVer=2.5.63&dict=spend%3D90%26result%3D0
`
	fmt.Println(queryStrParse(queryStr))
}
