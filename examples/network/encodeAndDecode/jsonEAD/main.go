package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	WithGF bool    `json"withGF"`
	Weight float64 `json:"-"`                // - 表示这个字段不会输出到json
	Height float64 `json:"height,omitempty"` // 表示如果字段为空，不会输出到json
}

func lazyJsonEncode() []byte {
	return []byte(`{
   "deptList": [
      {
         "id": 1,
         "name": "$技术部",
         "parentId": 0,
         "sortId": 1
      },
      {
         "id": 1,
         "name": "销售部",
         "parentId": 0,
         "sortId": 2
      }
   ]}`)
}
func main() {
	s := Student{
		Name:   "jack",
		Age:    23,
		WithGF: true,
		Weight: 53,
	}

	// json 编码
	jData, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jData)

	// json 解码
	student := Student{}
	err = json.Unmarshal(jData, &student)
	if err != nil {
		panic(err)
	}

	fmt.Println(student)

	// 直接 json 编码
	Data := lazyJsonEncode()
	fmt.Printf("%s", Data)
}
