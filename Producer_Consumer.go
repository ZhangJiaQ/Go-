package main

import "fmt"

/*
这个文件实现了一个多消费者与多生产者模型的简单示例
利用channel实现了生产者和消费者之间的数据和状态的同步

第一版遇到了
fatal error: all goroutines are asleep - deadlock!
这种所有线程或者进程都在等待资源释放的情况，我们便把它称之为死锁。
说明main中有goroutine等待从chan中读取消息
而chan已经关闭了
排查代码发现在消费者读取完成后，
忘了在doneCh的chan中放置消息，
而main函数最后还等待chan中的消息
所以导致了上述错误
*/

func Producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func Consumer(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch
		if ok {
			fmt.Printf("id %d take Value is %d\n",id,  value)
		} else {
			fmt.Printf("id %d is closed\n",id)
			break
		}
	}
	done <- true
}


func main() {
	ch := make(chan int, 3)
	cosNum := 2  // 消费者数目
	doneCh := make(chan bool, cosNum)

	for i:=0; i < cosNum; i++ {
		// 两个消费者
		go Consumer(i, ch, doneCh)
	}

	// 一个生产者
	go Producer(ch)

	for i:= 1; i <= cosNum; i++ {
		<- doneCh
	}

}