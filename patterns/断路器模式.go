package main

import (
	"time"
	"context"
)

/*
断路器是一个现代软件开发的设计模式。
用以侦测错误，并避免不断地触发相同的错误
（如维护时服务不可用、暂时性的系统问题或是未知的系统错误）。
*/

// 操作计数

/*
下面是一个简单的计数器
用于记录再一次循环中成功和失败的状态以及失败的时间和统计
*/

type State int

const (
	UnknownState State = iota
	FailureState
	SuccessState
)

type Counter interface {
	Count(State)
	ConsecutiveFailures() uint32
	LastActivity() time.Time
	Reset()
}



// 循环破坏者

/*
计数器在闭包内被Breaker调用来保持一个内部的操作计数器
如果circuit失败超过阈值，Breaker返回一个快速失败
过一会儿Breaker会重新记录请求
*/

func NewCounter() Counter {
	var i Counter
	return i
}

type Circuit func(context.Context) error

func Breaker(c Circuit, failureThreshold uint32) Circuit{

	cnt := NewCounter()

	return func(ctx context.Context) error {
		// pass
		if cnt.ConsecutiveFailures() >= failureThreshold {
			canRetry := func(cnt Counter) bool {
				backOffLevel := cnt.ConsecutiveFailures() - failureThreshold

				// 计算断路器何时恢复向服务传播请求
				shouldRetryAt := cnt.LastActivity().Add(time.Second * 2 << backOffLevel)

				return time.Now().After(shouldRetryAt)
			}

			if !canRetry(cnt) {
				// 快速失败而不是将请求传播到Circuit，因为自上次重试失败以来没有足够的时间
				return ErrIllegalRelease
			}
		}

		// 除非超过失败阈值，否则包装的服务会模仿旧的行为，并且在连续失败后会看到行为的差异

		if err := c(ctx); err != nil {
			cnt.Count(FailureState)
			return err
		}

		cnt.Count(SuccessState)
		return nil
	}


}
