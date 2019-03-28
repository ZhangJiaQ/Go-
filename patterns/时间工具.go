package main

import (
	"time"
	"log"
)

/*
有时候我们需要快速统计函数运行时间
则可以用一下方式
*/

func Duration(invocation time.Time, name string){
	elapsed := time.Since(invocation)
	log.Printf("%s lasted %s", name, elapsed)
}


func main() {
	defer Duration(time.Now(), "IntFactorial")

	time.Sleep(time.Second*2)

}