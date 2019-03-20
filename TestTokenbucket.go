package ch8

import (
	"time"
	"sync"
	"container/list"
	"fmt"
)

/*
令牌桶算法是网络流量整形（Traffic Shaping）和速率限制（Rate Limiting）中最常使用的一种算法。典型情况下，
令牌桶算法用来控制发送到网络上的数据的数目，并允许突发数据的发送。
令牌桶算法的基本过程如下：
假如用户配置的平均发送速率为r，则每隔1/r秒一个令牌被加入到桶中；
假设桶最多可以存发b个令牌。如果令牌到达时令牌桶已经满了，那么这个令牌会被丢弃；
当一个n个字节的数据包到达时，就从令牌桶中删除n个令牌，并且数据包被发送到网络；
如果令牌桶中少于n个令牌，那么不会删除令牌，并且认为这个数据包在流量限制之外；
算法允许最长b个字节的突发，但从长期运行结果看，数据包的速率被限制成常量r。对于在流量限制外的数据包可以以不同的方式处理：
它们可以被丢弃；
它们可以排放在队列中以便当令牌桶中累积了足够多的令牌时再传输；
它们可以继续发送，但需要做特殊标记，网络过载的时候将这些特殊标记的包丢弃。
*/

type TokenBucket struct {
	interval time.Duration    // 时间间隔
	ticker *time.Ticker       //  定时器
	cap int64    // 容量
	avail int64    // 可用数量
	tokenMutex *sync.Mutex  //用于令牌取出加入锁
	waitingQuqueMutex  *sync.Mutex  // 等待队列的锁
	waitingQuque *list.List    // 等待队列
}

type waitingJob struct {
	ch chan struct {}
	count int64
	need int64      // 需要令牌数量
	abandoned bool  // 超时弃用
}

func (tb *TokenBucket) addWaitingJob(w *waitingJob) {
	// 给队列末尾添加一个元素
	tb.waitingQuqueMutex.Lock()
	tb.waitingQuque.PushBack(w)
	tb.waitingQuqueMutex.Unlock()
}

func (tb *TokenBucket) getFrontWaitingJob() *list.Element {
	// 获取队列中先进的队首元素
	tb.waitingQuqueMutex.Lock()
	res := tb.waitingQuque.Front()
	tb.waitingQuqueMutex.Unlock()
	return res
}

func (tb *TokenBucket) removeWaitingJob(e *list.Element)  {
	// 删除队列中队首元素
	tb.waitingQuqueMutex.Lock()
	tb.waitingQuque.Remove(e)
	tb.waitingQuqueMutex.Unlock()
}


func (tb *TokenBucket) adjustDaemon(){
	var waitingJobNow *waitingJob
	for now := range tb.ticker.C{
		// 每隔一定时间段， 增加一个一个令牌
		var _ = now
		tb.tokenMutex.Lock() // 加入令牌加锁
		if tb.avail < tb.cap {
			tb.avail ++
		}
		// 查询等待队列中是否有等待取令牌的
		element := tb.getFrontWaitingJob()
		if element != nil {
			if waitingJobNow == nil || waitingJobNow.abandoned{
				waitingJobNow = element.Value.(*waitingJob)
				tb.removeWaitingJob(element)  // 取出后移除队首元素
			}
			if tb.avail >= waitingJobNow.need && !waitingJobNow.abandoned{
				// 可用大于 队列队首需要的
				waitingJobNow.ch <- struct{}{}  // 给channel加入值
				<-waitingJobNow.ch             //
				waitingJobNow = nil
			}
		}
		tb.tokenMutex.Unlock() // 加入令牌解锁
	}
}

func New(interval time.Duration, cap int64) *TokenBucket {
	if interval < 0 {
		panic(fmt.Sprintf("ratelimit: interval %v should > 0", interval))
	}

	if cap < 0{
		panic(fmt.Sprintf("ratelimit: capability %v should > 0", cap))
	}


	tb := &TokenBucket{
		interval:          interval,
		tokenMutex:        &sync.Mutex{},
		waitingQuqueMutex: &sync.Mutex{},
		waitingQuque:      list.New(),
		cap:               cap,
		avail:             cap,
		ticker:            time.NewTicker(interval),
	}
	go tb.adjustDaemon()  // 开始每隔一段时间添加令牌

	return tb
}

/*
不仅如此，我们还需要实现一下方法，用于令牌的取出
而且支持多goroutine
TryTake(count int64) bool： 尝试从桶中取出 n 个令牌。立刻返回，返回值表示该次取出是否成功。

Take(count int64)：尝试从桶中取出 n 个令牌，
若当前桶中的令牌数不足，则保持等待，直至桶内令牌数量达标然后取出。

TakeMaxDuration(count int64, max time.Duration) bool：尝试从桶中取出 n 个令牌，
若当前桶中的令牌数不足，则保持等待，直至桶内令牌数量达标然后取出。
不过设置了一个超时时间 max ，若超时，则不再等待立刻返回，返回值表示该次取出是否成功。

Wait(count int64)：保持等待直至桶内令牌数大于等于 n 。

WaitMaxDuration(count int64, max time.Duration) bool 保持等待直至桶内令牌数大于等于 n ，
但设置了一个超时时间 max 。
*/

func (tb *TokenBucket)checkNeed(need int64){
	if need <= 0 || need > tb.cap {
		panic(fmt.Sprintf("token-bucket: count %v should be less than bucket's"+
			" capablity %v", need, tb.cap))
	}
}

func (tb *TokenBucket)tryTake(need int64) bool{
	// 尝试从桶中取出 n 个令牌。立刻返回，返回值表示该次取出是否成功。
	tb.checkNeed(need)

	tb.tokenMutex.Lock()
	defer tb.tokenMutex.Unlock()
	if tb.avail >= need {
		tb.avail -= need
		return true
	}
	return false
}

/*
等待型取出
对于 Take(count int64) 和 TakeMaxDuration(count int64, max time.Duration) bool
这样的等待型取出（尝试），情况别就有所不同了：
由于这两个操作都是需要进行等待被通知，故原本的主动加锁检查共享资源的方案已不再适合。
由于可能存在多个正在等待的操作，为了避免混乱，我们需要有个先来后到，最早等待的操作，首先获取令牌。
*/


func (tb *TokenBucket)takeAndWait(need int64){
	// 试从桶中取出 n 个令牌，若当前桶中的令牌数不足，则保持等待，直至桶内令牌数量达标然后取出。
	if ok:= tb.tryTake(need); ok {
		return
	}
	w := &waitingJob{
		ch: make(chan struct{}),
		count: need,
	}
	tb.addWaitingJob(w) // 将w加入队列， 需要为队列加锁

	// 等待可以取出
	<-w.ch
	tb.avail -= need
	w.ch <- struct{}{} //chan取出
	close(w.ch)
}

func (tb *TokenBucket) waitAndTakeMaxDuration(need int64, max time.Duration) bool {
	if ok := tb.tryTake(need); ok{
		return true
	}

	w:= &waitingJob{
		ch: make(chan struct{}),
		count: need,
		abandoned: false,
	}
	defer close(w.ch)

	tb.addWaitingJob(w)

	select {
	case <-w.ch:
		tb.avail -= need
		w.ch <- struct{}{}
		return true

	case <-time.After(max):
		w.abandoned = true
		return false
	}
}

