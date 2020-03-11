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

```go
package map_test

import (
	"testing"
	"fmt"
)

type Pet struct {

}

func (p *Pet) Speak(){
	fmt.Print("Speaking")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("...", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak(){
	// 复用Pet的Speak方法
	d.p.Speak()
}

func (d *Dog)  SpeakTo(host string){
	// 复用Pet的SpeakTo方法
	d.p.SpeakTo(host)
}



func TestUInitMap(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Judge")
	// 复合只能用父类的方法
}


func TestAccessNotExistingKey(t *testing.T) {

}
```

匿名重载.

# 错误机制

1. 及早失败

```go
func GetFibonaccil(count int) (numbers []int, err error) {
	if count < 2 || count > 100 {
		return nil, errors.New("count should be bewteen 2 and 100")
	}
	fibList := []int{1,1}
	for i:=2; i<count; i++ {
		fibList = append(fibList, fibList[i-2]+ fibList[i-1])
	}
	return fibList, nil
}


func TestAccessNotExistingKey(t *testing.T) {
	if result, err:=GetFibonaccil(-1); err != nil{
		t.Log(err)
	}  else {
		t.Log(result)
	}
}
```

2. panic 用于不可以恢复得错误， 会执行defer，
3. os.Exit 退出时候不会输出调用栈信息，不会执行defer
4. recover panic后恢复(注意recover得设计实现)

# 构建一个可复用得包

package

1. 基本复用模块代码

   以首字母大写来表明可被包外代码访问

2. 代码得package

3. init方法   

4. Concurrent_map github



# 依赖管理



# 并发编程

### 协程机制

```Go
func TestAccessNotExistingKey(t *testing.T) {
	for i:= 10; i > 0; i-- {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond*500)
}
```



### 共享内存并发机制

```go
func TestShareMem(t *testing.T) {
   var mut sync.Mutex
   count := 0
   for i:= 0; i < 5000; i++ {
      go func() {
         mut.Lock() // 锁住
         count++
         mut.Unlock()  // 解锁
      }()
   }
   time.Sleep(1*time.Second)
   fmt.Println(count)
}

func TestWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup  // 等待
	count := 0
	for i:= 0; i < 5000; i++ {
		wg.Add(1) // 增加要等待的量
		go func() {
			mut.Lock()
			count++
			mut.Unlock()
			wg.Done() // 任务结束减少要等待的量
		}()
	}
	wg.Wait() // 所有任务完成才会向下执行
	fmt.Println(count)
}

```

### CSP并发机制

通过channel进行通讯，



### 多路选择和超时

select

### channel的关闭和广播



# Go语言典型并发任务

1. 单例
2. 





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












