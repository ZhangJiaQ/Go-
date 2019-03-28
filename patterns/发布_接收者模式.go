package main

import (
	"errors"
	"time"
)

/*
接收发布者模式是一种消息模式，
利用消息传递在两个不同的对相见进行传递

*/

type Message struct {
	// 消息内容
}


type Subscription struct {
	// 发布者
	ch  chan  Message
	Inbox chan Message
}

func (s *Subscription) Publish(msg Message) error {
	// 发布者的发布功能
	if _, ok := <-s.ch; !ok {
		return errors.New("Topic has been closed")
	}
	// 写入消息到channel内
	s.ch <- msg
	return nil
}

type User struct {
	ID uint64
	Name string
}

type Session struct {
	User User
	Timestamp time.Time
}

type Topic struct {
	Subscribers []Session
	MessageHistory []Message
}

func (t *Topic) Subscribe(uid unit64) (Subscription, error) {
	// 如果uid是首次注册，则生成session否则就获取session

	// 加入session到Topic和MessageHistory内

	// 创建一个subscription
}

func (t *Topic) Unsubscribe(Subscription) error {
	// 执行取消订阅
}

func (t *Topic) Delete() error {
	// 执行删除
}

