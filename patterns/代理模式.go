package main

import (
	"fmt"
)

/*
代理模式（英语：Proxy Pattern）是程序设计中的一种设计模式。
所谓的代理者是指一个类别可以作为其它东西的接口。
代理者可以作任何东西的接口：网上连接、存储器中的大对象、文件或其它昂贵或无法复制的资源。
*/

// 要使用代理和对象，他们必须实现相同的方法
type IOMyObject interface {
	MyObjectDo(action string)
}

//
type MyObject struct {
	action string
}

// ObjDo实现了IObject接口的所有逻辑
func (obj *MyObject) MyObjectDo(action string) {
	fmt.Printf("I can, %s", action)
}


type ProxyObject struct {
	myobject *MyObject
}

// ProxyObject在真实对象操作发生之前实现了对之前对象的拦截并操作
func (p *ProxyObject) MyObjectDo(action string) {
	if p.myobject == nil {
		p.myobject = new(MyObject)
	}
	if action == "Run" {
		p.myobject.MyObjectDo(action)
	}
}