package main

import (
	"errors"
	"time"
)

var (
	ErrNoTickets = errors.New("semaphores: could not aquire semaphore")
	ErrIllegalRelease = errors.New("semaphores: can't release the semaphore without acquiring it first")
)

// pass
type SeamphoreInterface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem chan struct{}
	timeout time.Duration
}

func (s *implementation)Acquire() error {
	// 实现信号量的获取方法
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation)Release() error {
	// 实现信号量的释放方法
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
	return nil
}

func SeamphoreNew(tickets int, timeout time.Duration) SeamphoreInterface{
	return &implementation{
		sem: make(chan struct{}, tickets),
		timeout: timeout,
	}
}


func main() {
	tickets, timeout := 1, 3*time.Second
	s := SeamphoreNew(tickets, timeout)
	if err := s.Acquire(); err != nil {
		panic(err)
	}

	// 业务逻辑

	if err := s.Release(); err != nil {
		panic(err)
	}
}

