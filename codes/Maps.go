package main

import "fmt"

func main() {
	m := make(map[string]int)

	//往map中添加键-值
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	//将k1的值赋给v1
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//内置函数len可以返回map长度
	fmt.Println("len:", len(m))

	//内置函数delete可以删除map中的键-值对
	delete(m, "k2")
	fmt.Println("map:", m)

	//如果map中值不存在，返回0,flase    如果存在返回map[key],true
	x, y := m["k1"]
	fmt.Println(x, y)

	//初始化和赋值合为一步
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)
}
