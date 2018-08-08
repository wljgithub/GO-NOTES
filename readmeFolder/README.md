# purpose
将会记录自己学习go语言的笔记


- [语法概要](#语法概要)
    - []()
    - [关键字](#关键字)
    - [预定义标识符](#预定义标识符)
    - [操作符](#操作符)
    - [表达式](#表达式)

    - []()

- [基础数据类型](#基础数据类型)
    - [变量](#变量)
    - [变量默认值](#默认值)
    - [类型转换](#类型转换)
    - [常量](#常量)
    - []()

- [复合数据类型](#复合数据类型)
    - [数组](#数组)
    - [切片](#切片)
    - [Map](#Map)
    - [Struct](#Struct)
    - [自定义类型](#自定义类型)
    - [指针](#指针)


- [流程控制](#流程控制)
    - [If](#If)
    - [Switch](#Switch)
    - [For](#For)
    - []()

- []()
- []()

	
<details>
  <summary># go环境搭建</summary>
- [下载window64的go](https://golang.org/doc/install?download=go1.10.3.windows-amd64.msi)

- [在sublimetext上配置go环境](https://blog.csdn.net/u012855229/article/details/72642735)
</details>


## 语法概要

go语法元素由5类内容组成：


- 标识符(identifier)
- 关键字(keyword)
- 字面量(literal)
- 操作符(operator)
- 分隔符(delimiter)

&nbsp;
### 标识符

标识符指的是变量名、函数名、自定义的接口


go语言中，任何标识符(变量、常量、函数、自定义类型等)都应该满足以下规律：

- 以字符或下划线开头
- 不能与go关键字冲突

&nbsp;
除此之外，go还有36个预定义的标识符
### 预定义标识符

```
内置函数：append、 cap、 close、 complex、 copy、 delete、 imag、 len、 make、 panic、 print、 println、 real、 recover、 new

常量：true、 false、 iota

接口类型：error	nil

基础数据类型：
int、 int8、 int16、 int32、 int64、 uint、 uint8、 uint16、 uint32、 uint64、 uintptr、
float32、 float64、 complex64、 complex128、 string、 bool、	 byte、
```

&nbsp;
### 关键字

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
### 操作符

操作符分算数操作符、比较操作符、逻辑操作符、地址操作符、接受操作符

&nbsp;
### 操作符优先级
。。待续

&nbsp;
### 表达式

选择表达式、索引表达式、切片表达式、类型断言、调用表达式

&nbsp;

## 基础数据类型
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
### 变量

示例

### 默认值

没有明确初始值的变量，go会赋予他们默认值

| 数据类型 | 默认值 |
|----------|--------|
| int      | 0      |
| float64  | 0      |
| bool     | false  |
|          string| "" (空字符串)       |


<details>
  <summary>示例代码：</summary>

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

```
输出
```go
0 0 false ""
```
</details>



### 类型转换

Go 在不同类型的项之间赋值时需要显式转换

<details>
  <summary>示例代码：</summary>

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

```
输出：
```go
3 4 5
```
</details>




### 常量

- 常量不能用  := 语法声明，只能用 const
- 数值常量时高精度的值，数值常量是高精度的 值。

一个未指定类型的常量由上下文来决定其类型。

再尝试一下输出 needInt(Big) 吧。

（int 类型最大可以存储一个 64 位的整数，有时会更小。）

（int 可以存放最大64位的整数，根据平台不同有时会更少。）


&nbsp;
## 复合数据类型

go 的基本数据类型都完整地确定了类型的方方面面，而高级数据类型服务于用户自己定义的数据类型，你可以制定元素类型和元素长度来定义一个数组，也可以指定键类型和元素类型来声明一个字典类型，go的高级数据类型相当于自定义数据类型的模板和制作工具

&nbsp;
### 数组
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



### 切片

每个数组的大小都是固定的，而切片则为数组元素提供动态大小得、灵活的视角

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
- 切片并不存储任何数据，它只是描述底层数组中的一段
- 更改切片的元素则会修改数组中对应的元素，和共享底层数组切片的元素



<details>
  <summary>示例代码：</summary>

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```
输出：
```go
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
```
</details>



切片的默认：
在进行切片时，你可以利用它的默认行为来忽略上下界。

切片下界的默认值为 0，上界则是该切片的长度

<details>
  <summary>示例代码：</summary>

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```
输出：
```go
[3 5 7]
[3 5]
[5]
```
</details>

- 切片的零值时nil

 
<h3 id="Map">Map</h3>
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


数组和切片都只能存储同一种类型的数据，如果想存储不同类型的数据需要用到结构体



```go
语法： type 名 struct

如：
type student struct{
	age int
	name string
}
```

<h3 id="Struct">Struct</h3>
示例：

<details>
  <summary>示例代码：</summary>

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```
</details>






需要注意的是：

- 结构体中的字段名不能相同
- 当结构体的字段名为小写字母开头时不对外开放，即其他包无法直接使用该字段名
- 结构体字段使用点号来访问


切片中可以内嵌结构体：



### 自定义类型
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



<details>
  <summary>示例代码：</summary>

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
</details>

示例：

- 需要注意的是，虽然可以像操作原类型一样操作自定义的类型，但他们已经是不同类型了，在将参数传入函数中时需要注意类型转换，否则会编译错误


### 指针

& 操作符会生成一个指向操作数的指针，本质上是一个16进制的内存地址

 *操作符表示指针指向的底层值


<details>
  <summary>示例代码：</summary>

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```
输出：
```go
42
21
73
```
</details>




结构体指针：

如果有一个指向结构体的指针p，那么可以通过 *p.X 来访问X，不过在go中，允许缩写成p.X


<details>
  <summary>示例代码：</summary>
  
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```
输出：
```go
{1000000000 2}
```
</details>

## 流程控制



### 
<h3 id="If">if</h3>

Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
```go
package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
```
同 for 一样， if 语句可以在条件表达式前执行一个简单的语句,该语句声明变量的***作用域*** 仅在 if 之内(包含else字句)
```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```
输出：
```go
9 20
```
<h3 id="Switch">Switch</h3>
switch 是编写一连串 if - else 语句的简便方法，

Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数
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

	//没有条件的switch

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

```

<h3 id="For">For</h3>

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

continue则跳到下一次循环
```go

for i:=0;i<5,i++{
	if i%2 == 0{
		continue
	}
	fmt.Println(i)
}
```
#### 函数

函数也是值，它可以像其他值一样传递
函数值可以用作函数的参数和返回值
```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```
输出：
```go
13
5
81
```
函数的闭包：

一般情况下变量脱离了它的作用域，就会消失了，但是通过闭包可以将其绑定在一个函数上，
闭包本质上就是一个嵌套函数，内层函数引用外层函数的变量值
```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```
输出
```go
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
```
#### 方法

#### 接口


&nbsp;




### 异常处理

#### defer

defer 语句会将函数推迟到外层函数返回之后执行
```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

```
输出：
```
hello
world
```

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
- 指针到底时用来干什么的，修改变量的值？直接改不就行了吗，干嘛要多引出一个指针出来

# 参考


- [Go 零基础编程入门教程](https://songjiayang.gitbooks.io/go-basic-courses/content/ch3/identifiers.html)
- [go知识蓝图](https://www.processon.com/view/link/5a9ba4c8e4b0a9d22eb3bdf0)

