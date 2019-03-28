package main

/*
当对象间存在一对多关系时，则使用观察者模式（Observer Pattern）。
比如，当一个对象被修改时，则会自动通知它的依赖对象。观察者模式属于行为型模式。
*/

type (
	// Event定义了一个“事件”发生时的迹象
	Event struct {
		// data在这次讨论中是一个简单的Int类型变量，实际上类型应该取决于你的应用
		Data int64
	}

	// Observer为这个示例提供了一个标准的interface， 用于列出发生的“事件”
	Observer interface {
		OnNotify(Event)
	}

	Notifier interface {
		// Register允许一个示例注册他用于监听观察事件
		Register(Observer)
		// Register允许一个示例移除他用于监听观察事件
		Unregister(Observer)
		// 提醒注册过的对象事件发生
		Notify(Event)
	}


)


