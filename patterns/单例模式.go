package main

import (
	"sync"
	"fmt"
)

/*
Singleton创建设计模式将类型的实例化限制为单个对象。
单例模式表示全局状态，大多数情况下会降低可测试性。
*/


type singleton map[string]string

var (
	once sync.Once
	instance singleton
)

func New() singleton {
	once.Do(func(){
		// 初始化一个map
		instance = make(singleton)
	})
	return instance
}


func main() {
	s := New()

	s["this"] = "that"

	s2 := New()

	fmt.Println("this is ", s2["this"])
}