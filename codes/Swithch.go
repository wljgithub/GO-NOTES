package main

import (
	"fmt"
	"time"
)

func main() {

	// 一般用法
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//可用逗号分隔多个case的表达式
	//我在这里也用上default

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Println("It it the weekday")
	}

	//也可以把switch当if-else用

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

	//也可以通过switch来判断值的类型
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		default:
			// fmt.Println("Don't know type %T\n", t)
			fmt.Printf("Don't know type %T\n", t)

		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hty")
}
