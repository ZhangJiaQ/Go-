package main

import "fmt"

/*
通过对象池可以减少对象创建的消耗
在对象初始化比对象维护消耗更多的情况下，对象池模式很有用。
由于预先初始化对象，它对性能有积极影响。
当然，我们也可以直接用sync.pool内建对象池
*/

type Object struct {
	name string
	age int
}

func (Object)Do(){
	fmt.Println("Do")
}

type Pool chan *Object


func New_(total int) Pool {
	// 创建一个包含对象的chan
	p := make(Pool, total)

	// 创建对象进入chan(对象池)
	for i:= 0; i<total; i++ {
		p <- new(Object)
	}

	return p
}

func main(){
	p := New_(2)
	select {
		case obj := <-p:
			// 取出对象 使用对象方法
			obj.Do()
			// 放回对象
			p <- obj

		default:
			return

	}
}
