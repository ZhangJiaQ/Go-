package main

import "sync"

/*
Merge函数通过为每个输入channel启动goroutine
将一个channel列表转换为单个channel，
该channel将值复制到唯一的输出channel。
一旦启动了所有输出goroutine，
就会启动Merge a goroutine来关闭主channel。
*/

func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	// send 为一个函数对象
	// 接收一个channel 遍历出channel中所有的值，然后返回，并告诉waitgroup已经完成
	// 类似生成器这种
	// 知道channel关闭
	send := func(c <-chan int) {
		for n := range c {
			out <-n
		}
		wg.Done()
	}
	// 增加cs长度个WaitGroup, 调用输入到主channel内
	wg.Add(len(cs))

	for _, c := range cs {
		go send(c)
	}

	// 完成后关闭channel
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}



/*
Split函数将一个channel的数据分给多个channel
*/

func Split(ch chan int, n int) [] chan int {
	cs := make([]chan int, n)
	// 多个chan加入到chan列表中
	for i:= 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	// 设置分解函数
	distributeToChannels := func(ch chan int , cs []chan int) {
		// 退出后关闭所有channel
		defer func(cs []chan int) {
			for _, c := range cs {
				close(c)
			}
		} (cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <- ch:
					if !ok{
						return
					}
					c <- val
				}
			}
		}

	}
	// 调用分解函数
	go distributeToChannels(ch, cs)

	return cs
}