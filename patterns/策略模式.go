package main

/*
在策略模式（Strategy Pattern）中，
一个类的行为或其算法可以在运行时更改。这种类型的设计模式属于行为型模式。

在策略模式中，我们创建表示各种策略的对象和一个行为随着策略对象改变而改变的 context 对象。
策略对象改变 context 对象的执行算法
意图：定义一系列的算法,把它们一个个封装起来, 并且使它们可相互替换。
主要解决：在有多种算法相似的情况下，使用 if...else 所带来的复杂和难以维护。
何时使用：一个系统有许多许多类，而区分它们的只是他们直接的行为。
如何解决：将这些算法封装成一个一个的类，任意地替换。
关键代码：实现同一个接口。
*/

// 策略定义
// 策略实现一个方法 方法会调用其接口中的Apply方法
// 我们可以在接下来的策略中实现不用的Apply方法，进行各种不同的策略调用
type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftvalue, rightvalue int) int {
	return o.Operator.Apply(leftvalue, rightvalue)
}

// 加法策略
type Addition struct{
}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

// 乘法策略
type Multiplication struct{
}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}


func main() {
	Operator := &Operation{}

	// 加法策略
	addMethod := &Addition{}
	Operator.Operator = addMethod
	add_number := Operator.Operate(52,2)
	print(add_number)

	// 乘法策略
	multiplicationMethod := &Multiplication{}
	Operator.Operator = multiplicationMethod
	multiplie_number := Operator.Operate(52,2)
	print(multiplie_number)
}

