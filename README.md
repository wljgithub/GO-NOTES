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

	程序声明： import package

	程序实体声明和定义： chan、const、func、interface、map、struct、type、var

	程序流程控制： go、select、break、case、continue、default、defer、else、fallthrough、for、goto、if、range、return、switch

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
	bool      the set of boolean (true, false)

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

	byte      alias for uint8
	rune      alias for int32
	uint      either 32 or 64 bits
	int      same size as uint
	uintptr      an unsigned integer large enough to store the uninterpreted bits of a pointer value

	string      the set of string value (eg: "hi")

&nbsp;
### 高级数据类型

go 的基本数据类型都完整地确定了类型的方方面面，而高级数据类型服务于用户自己定义的数据类型，你可以制定元素类型和元素长度来定义一个数组，也可以指定键类型和元素类型来声明一个字典类型，go的高级数据类型相当于自定义数据类型的模板和制作工具

&nbsp;
#### 数组

	var ipv4 = [4]unit8{192.168.0.1}

- 在函数中，var关键字可省略，赋值符号由 = 改为 :=
- 数组一旦声明，长度不可改变
- 数组长度也是数据类型的一部分，即使数组中元素类型相同，但长度不同，它们都不属于同种类型 
- 内置函数len和cap可以作用于数组
- 数组长度是不可变的，可以完全避免耗时耗力的内存二次分配操作



#### 切片

#### Map

#### 自定义类型

#### 结构体

#### 函数

#### 方法

#### 接口

&nbsp;
## go语言变量
&nbsp;
### 变量声明有三种方式

第一种，指定变量类型，声明后若不复制，使用默认值

	var var_name v_type
	v_name = value


第二种，根据值自行判断变量类型

	var v_name = value

第三种，省略var，注意:=左侧的变量不能是已经声明过的，否则导致编译错误

	v_name:= value
&nbsp;
###  多变量声明

也是对应的上面三种

一：

	var name1,name2,name3 type
	name1,name2,name3 = v1,v2,v3

二：

	var name1,name2,name3 = v1,v2,v3
	
三：

	name1,name2,name3 =: v1,v2,v3

## 流程控制

### 分支循环

#### if

#### switch

#### for


### 异常处理

# go代码片段


- [hello world!](https://github.com/wljgithub/go-/blob/master/codes/hello.go)
    + 每一个独立的go程序，都必定包含一个packge main
    + 

-[]()


# 参考


- [Go 零基础编程入门教程](https://songjiayang.gitbooks.io/go-basic-courses/content/ch3/identifiers.html)
- [go知识蓝图](https://www.processon.com/view/link/5a9ba4c8e4b0a9d22eb3bdf0)

