package main

import "testing"

func Test_Division(t *testing.T)  {
	if val,err:=Division(6,3);val ==2 &&err==nil {
		t.Log("Division函数 测试通过","值为：",val)
	}else {
		t.Error("Division函数 测试失败")
	}
}

func Test_Division_2(t *testing.T) {
	t.Error("就是不通过")
}