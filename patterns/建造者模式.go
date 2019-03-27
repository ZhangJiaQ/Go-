package main

/*
 定义：  建造者设计模式是一个构造复杂对象的设计模式，在一个软件系统中，可能会面临创建一个复杂对象的工作，
如果使用单一的方法或者单一的对象来创建会比较烦琐，当所创建复杂对象发生改变时，整个系统就可能面临剧烈的变化。
这时就需要我们将这个复杂对象的创建过程分解成若干部分，各个子部分用一定的算法构成。
但是，子部分可能会经常发生变化，如何能保证整体创建工作的稳定性呢？
这就需要建造者模式，建造者模式把复杂对象的创建与表示分离，使得同样的构建过程可以创建不同的表示。
*/


type Speed float64

const (
	MPH Speed = 1
	KPH 	  = 1.60934
)

type Color string

const (
	BlueColor Color = "blue"
	GreenColor      = "green"
	RedColor        = "red"
)

type Wheels string

const (
	SportWheels Wheels = "sports"
	SteelWheels 	    = "Steel"
)

type Builder interface {
	 Color(Color) Builder
	 Wheels(Wheels) Builder
	 TopSpeed(Speed) Builder
	 Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

func main(){
	assembly := Builder.Color(BlueColor)

	familyCar := assembly.Wheels(SportWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()
	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()
}