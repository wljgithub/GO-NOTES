package main

import (
	"encoding/json"
	"fmt"
)

// 编码json格式的数据，用byte[]即可，注意json要用{}包着
func encodeJson(p ...interface{}) (js []byte) {
	return []byte(fmt.Sprintf(`{
   "deptList": [
      {
         "id": %d,
         "name": "%s",
         "parentId": %d,
         "sortId": %d}]}`, p[0], p[1], p[2], p[3]))
}

func decodeJson(js []byte, contaiter *map[string]interface{}) {
	json.Unmarshal(js, &contaiter)
}
func main() {

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	json.Unmarshal(byt, &dat)
	fmt.Println(dat)

	// 编码json 格式的数据
	s := []byte(`{
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
   ]
}`)
	// 解码json任意格式的数据
	var unknowDataType map[string]interface{}
	json.Unmarshal(s, &unknowDataType)
	fmt.Println(unknowDataType)

	decodeJson(s, &unknowDataType)
	fmt.Println(unknowDataType)

	//fmt.Println(string(encodeJson(0,"技术部",1,1)))
}
