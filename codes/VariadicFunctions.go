package main

import "fmt"

// 任意个参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// 可以将整个数组传入函数
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
