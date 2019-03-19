# 变量与常量



# 数据类型



# 运算符





# 条件与循环



# 数组与切片



# Map



# 字符串



# Go函数



# 可变长参数, defer

### 可变长参数

支持可变长参数列表的函数可以支持任意个传入参数，比如fmt.Println函数就是一个支持可变长参数列表的函数。

```go
func NumSum(numbers ...int) int {
   result := 0
   for _, num := range numbers {
      result += num
   }
   return result
}

func TestSliceInit(t *testing.T) {
   numbers := NumSum(1,2,3,4,5)
   println(numbers)
}
```

defer



# 行为的定义或实现



# Go语言interface

### 什么是interface

简单的说，interface是一组method的组合，我们通过interface来定义对象的一组行为。

也就是定义对象之间交互协议

### interface类型

interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

```go
package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}
func (h Human) SayHi(){
	fmt.Printf("Hi, I am %s. My phone is %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("LALALALA%s", lyrics)
}

func (e Employee) SayHi() {
	// Employee 重载Human SayHi 的方法
	fmt.Printf("I am %s, I work at %s\n", e.name, e.company)
}

type Men interface {
	// Interface Men被Human Student Employee实现
	SayHi()
	Sing(lyrics string)
}

func main(){
	// interface就是一组抽象方法的集合，
	// 它必须由其他非interface类型实现，
	// 而不能自我实现
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	//Tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}
	// 定义Men类型的变量i
	var i Men
	// i能储存Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("Student")

	// i也能储存Employee
	i = sam
	fmt.Println("This is sam, a Employee:")
	i.SayHi()
	i.Sing("Employee")

	// i也能定义slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x {
		value.SayHi()
	}

}
```

### 空interface

类似C语言的void*,可以春促任意类型的数值，对描述起不到任何作用。

```go
// 定义a为空接口
var a interface{}
var i int = 5
s := "Hello world"
// a可以存储任意类型的数值
a = i
a = s
```

### interface函数参数

pass

### interface变量存储的类型

我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：

1. Comma-ok断言   value, ok = element.(T)

   如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false，

2. if swich 



# 扩展与复用







# 如何用Go操作JSON

### 结构化数据

#### 普通Json

如同{"name": "Judge","age": "1"}这种Json字符串，我们在操作的时候需要先定义字符串

```go
type Student struct {
	Name string
	Age string
}

func TestSliceInit(t *testing.T) {
	ourStudentJson := `{"name": "Judge","age": "1"}`
	var stu Student
	json.Unmarshal([]byte(ourStudentJson), &stu)
	fmt.Printf("name: %s", stu.Age)
}

```

#### json数组

```Go
type Student struct {
	Name string
	Age string
}

func TestSliceInit(t *testing.T) {
	ourStudentJson := `[{"name": "Judge","age": "1"}, {"name": "Michael","age": "2"}]`
	var stu []Student
	json.Unmarshal([]byte(ourStudentJson), &stu)
	for _, st := range stu {
		fmt.Printf("Name is %s, Age is %s\n", st.Name, st.Age)
	}
}
```

#### 嵌套对象



### 非结构化数据

# Go单元测试



# GORM操作数据库

引入包

```go
import (
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
)
```



数据库连接

```go
func main() {
   db, err := gorm.Open("mysql", "root:123456@/tbkt_active?charset=utf8&parseTime=True&loc=Local")
   defer db.Close()
   if err != nil {
      fmt.Print("Error", err)
   }
}
```












