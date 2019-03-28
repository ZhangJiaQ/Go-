package main

import "fmt"

/*
Count 直接返回了一个channel c,
这个channel c使用一个旁路（goroutine)在一旁一直尝试往里面放数值。如果不能放，
这个 goroutine阻塞等待。在main函数中，使用range从channel c中取值。

*/

func Count(start int, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i:= start; i<=end; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}

func main(){
	for i:= range Count(1, 99) {
		fmt.Println(i)
	}
}