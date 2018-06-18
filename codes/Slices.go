package main

import "fmt"

func main() {

	// 用内置函数make创建一个切片
	s := make([]string, 3)
	fmt.Println("emp:", s)

	//和操控数组一样操控切片
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 内置函数len可返回切片长度
	fmt.Println("len:", len(s))

	//通过内置函数append来创建一个切片，并添加新的元素
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//通过内置函数copy复制一个切片
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//切片这种数据结构同样“切片操作”
	l := s[2:5] //-->s2,s3,s4
	fmt.Println("sl1:", l)

	l = s[:5] //-->s0,s1,s2,s3,s4
	fmt.Println("sl2:", l)

	l = s[2:] //--> s2,s3,s4,s5
	fmt.Println("sl3:", l)

	// 初始化和声明切片可在一步完成
	t := []string{"g", "h", "l"}
	fmt.Println("dlc", t)

	// 切片还能组成多维数据结构，与数组不同的是内层切片长度可以改变
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)

}
