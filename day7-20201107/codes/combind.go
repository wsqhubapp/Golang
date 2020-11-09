package main

// 组合是为了代码的复用    继承
// 命名组合

// 发送接口

type Connection interface {
	Open() error
	Send(string) error
	Recive() error
	Close() error
}
