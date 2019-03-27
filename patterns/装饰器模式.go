package main

import "log"

/*
装饰器为
*/

func LogDecorate(fn func(int) int) func(int) int {
	return func(i int) int {
		log.Println("Starting with number ", i)
		result := fn(i)
		log.Println("End with number ", result)
		return result
	}
}

func Double(n int) int {
	return n * 2
}

func main(){
	f := LogDecorate(Double)
	f(5)
}