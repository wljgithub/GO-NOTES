package jsonEAD

import "encoding/json"

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	WithGF bool   `json"withGF"`
}

func jsonEncode(student Student) ([]byte, error) {
	return json.Marshal(student)
}

func jsonDecode(js []byte)(student *Student) {
	json.Unmarshal(js,&student)
}
