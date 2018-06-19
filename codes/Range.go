package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 迭代数组
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 迭代map中的键-值
	kvs := map[string]string{"a": "apple", "b": "bnana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 只迭代map中的键
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 也可以用range函数来迭代字符串，i为索引值，c为unicode对应的编码数
	for i, c := range "巫凌健" {
		fmt.Println(i, c)
	}
}
