# purpose
将会记录自己学习go语言的笔记

# go环境搭建

- [下载window64的go](https://golang.org/doc/install?download=go1.10.3.windows-amd64.msi)

- [在sublimetext上配置go环境](https://blog.csdn.net/u012855229/article/details/72642735)

# go语法

go语法元素由5类内容组成：


- 标识符(identifier)
- 关键字(keyword)
- 字面量(literal)
- 操作符(operator)
- 分隔符(delimiter)

&nbsp;
## 标识符

标识符指的是变量名、函数名、自定义的接口


go语言中，任何标识符(变量、常量、函数、自定义类型等)都应该满足以下规律：

- 以字符或下划线开头
- 不能与go关键字冲突

&nbsp;
除此之外，go还有36个预定义的标识符
### 36个预定义标识符

	append	bool	byte	cap	close	complex	complex64	complex128	uint16
	copy	false	float32	float64	imag	int	int8	int16	uint32
	int32	int64	iota	len	make	new	nil	panic	uint64
	print	println	real	recover	string	true	uint	uint8	uintptr


&nbsp;
## 关键字

go中有25个关键字，分三类
```
程序声明： import package

程序实体声明和定义： chan、const、func、interface、map、struct、type、var

程序流程控制： go、select、break、case、continue、default、defer、else、fallthrough、for、goto、if、range、return、switch
```
&nbsp;
## 字面量

字面量就是值的一种标记法，用于表示某个值属于什么数据类型的

也分三类：

- 表示基础数据类型值的各种字面量，如浮点值 12E-3
- 用于构造各种自定义的复合数据类型的类型字面量，
	
如定义了一个名为Name的结构体类型：

	type Name struct {
		Forename string
		Surname string
	}

- 用于表示复合数据类型的值的符合字面量 (这个看书的时候没太懂)

&nbsp;
## 操作符

操作符分算数操作符、比较操作符、逻辑操作符、地址操作符、接受操作符

&nbsp;
### 操作符优先级
。。待续

&nbsp;
## 表达式

选择表达式、索引表达式、切片表达式、类型断言、调用表达式

&nbsp;
## go数据类型

### 基础数据类型
```go
布尔类型：
bool      the set of boolean (true, false)

数字类型：
uint8      the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8      the set of all signed  8-bit integers (-128 to 127)
int16      the set of all signed 16-bit integers (-32768 to 32767)
int32      the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64      the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32      the set of all IEEE-754 32-bit floating-point numbers
float64      the set of all IEEE-754 64-bit floating-point numbers

complex64      the set of all complex numbers with float32 real and imaginary parts
complex128      the set of all complex numbers with float64 real and imaginary parts
int      same size as uint
uint      either 32 or 64 bits
uintptr      an unsigned integer large enough to store the uninterpreted bits of a pointer value

字符串：
string      the set of string value (eg: "hi")

特殊类型：
byte      uint8的别名，用于表示二进制数据的bytes
rune      int32 别名，用于表示一个符号

```
&nbsp;
### 高级数据类型

go 的基本数据类型都完整地确定了类型的方方面面，而高级数据类型服务于用户自己定义的数据类型，你可以制定元素类型和元素长度来定义一个数组，也可以指定键类型和元素类型来声明一个字典类型，go的高级数据类型相当于自定义数据类型的模板和制作工具

&nbsp;
#### 数组
```go
//int数组初始化后，默认值为0
var a [5]int
fmt.Println("emp:", a)

//可通过索引访问或改变数组的值(索引值从0开始)
a[4] = 100
fmt.Println("set:", a)
fmt.Println("get:", a[4])

//内置函数len可返回数组长度
fmt.Println("len:", len(a))

//数组初始化和赋值可同时进行
b := [5]int{1, 2, 3, 4, 5}
fmt.Println("dlc", b)
```



- 在函数中，var关键字可省略，赋值符号由 = 改为 :=
- 数组一旦声明，长度不可改变
- 数组长度也是数据类型的一部分，即使数组中元素类型相同，但长度不同，它们都不属于同种类型 
- 内置函数len和cap可以作用于数组
- 数组长度是不可变的，可以完全避免耗时耗力的内存二次分配操作



#### 切片

```go
t := make([]string,3)
t[0] = "a"
t[1] = "b"
t[2] = "c"

//也可以初始化并声明一个切片合为一步
t := []string{"b", "b", "c"}

// 通过索引访问或修改切片的值
s[0] = "a"
fmt.Println(s[1])

//内置函数len可以获取切片的长度
len(t)

//内置函数append可以创建一个切片并添加新的元素
s = append(s,"d")

//内置函数copy可以复制一个新的切片
y := make([]string,len(t))
copy(y,t)
```


#### Map
Map 是go中的内置关联数据结构，有时在其他编程语言中被称为字典。
```go
//通过内置函数make创建一个map
m := make(map[string]int)
m["k1"] = 7
m["k2"] = 13

//初始化并赋值合为一步
n := map[string]int{"foo":1,"bar":2}

//map的长度
len(n)

//删除map中的元素
delete(m,"k2")

//如果map中值不存在，返回0,flase    如果存在返回map[key],true
x, y := m["k1"]
```


#### 自定义类型
go 中支持用户自定义一些特殊的数据结构
```go
语法： type 名字 类型

声明一个：
type NAME string

声明多个：
type(
	B0 = int8
	B1 = int16
	B2 = int32
	B3 = int64
)

```
示例：
```go
package main

import "fmt"

type City string

func main() {
	//注意需要用括号包起来
    city := City("上海")
    fmt.Println(city)
}
```
- 需要注意的是，虽然可以像操作原类型一样操作自定义的类型，但他们已经是不同类型了，在将参数传入函数中时需要注意类型转换，否则会编译错误

#### 结构体
数组和切片都只能存储同一种类型的数据，如果想存储不同类型的数据需要用到结构体

```go
语法： type 名 struct

如：
type student struct{
	age int
	name string
}
```
示例：
```go
package main

import "fmt"

type Student struct {
    Age     int
    Name    string
}

func main() {
    stu := Student{
        Age:     18,
        Name:    "name",
    }
    fmt.Println(stu)

    // 在赋值的时候，字段名可以忽略
    fmt.Println(Student{20, "new name"})
}
```
输出：
```go
{18 name}
{20 new name}
```
需要注意的是：

- 结构体中的字段名不能相同
- 当结构体的字段名为小写字母开头时不对外开放，即其他包无法直接使用该字段名
#### 函数

#### 方法

#### 接口

&nbsp;



## 流程控制

### 分支循环

#### if
```go
package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has mutiple digits")
	}
}

```

#### switch
```go
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

```

#### for

for 是go语言唯一的循环结构
```go
语法格式： for initial;contion;after{

	...
}
```
例如：
```go
for i:=0;i<100;i++{

	fmt.Println(i)
}
```

也可以使用break和continue关键字控制循环跳转

```go	
for {
	fmt.Println("loop")
	break
}
```

continue则调到下一个循环
```go

for i:=0;i<5,i++{
	if i%2 == 0{
		continue
	}
	fmt.Println(i)
}
```
### 异常处理

#### defer

#### panic

#### recover


# go代码片段



- [Hello World!](https://github.com/wljgithub/go-/blob/master/codes/hello.go)
    + 每一个独立的go程序，都必定包含一个packge main
    + 
- [Values](https://github.com/wljgithub/GO-NOTES/blob/master/codes/Values.go)
- [Variables](https://github.com/wljgithub/GO-NOTES/blob/master/codes/Variables.go)
- [Constants](https://github.com/wljgithub/GO-NOTES/blob/master/codes/Constants.go)
- [For](https://github.com/wljgithub/GO-NOTES/blob/master/codes/For.go)
- [If/Else](https://github.com/wljgithub/GO-NOTES/blob/master/codes/If-else.go)
- [Switch](https://github.com/wljgithub/GO-NOTES/blob/master/codes/Swithch.go)
- [Arrays](https://github.com/wljgithub/GO-NOTES/blob/master/codes/Arrays.go)
- [Slices](https://github.com/wljgithub/GO-NOTES/blob/master/codes/Slices.go)
- [Maps]()


# tips

- go中的字符串只能用双引号"",不可以用单引号
- go中引进的包如果不使用，则会编译错误
- if-else判断是，else不要换行，不然会编译错误
- go中声明的变量得如果不调用，则会编译错误，如果想声明但不调用，可以用_下划线作变量名

# FAQ

- go中的Print、Println有什么区别 
- 奇怪，go中也有栈溢出的说法吗?我写个递归函数，递归700层的时候值直接为0了

# 参考


- [Go 零基础编程入门教程](https://songjiayang.gitbooks.io/go-basic-courses/content/ch3/identifiers.html)
- [go知识蓝图](https://www.processon.com/view/link/5a9ba4c8e4b0a9d22eb3bdf0)

