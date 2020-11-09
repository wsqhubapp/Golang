// 1 接口 定义
package main

import (
	"fmt"
)

/*
需求：监控产生报警需要发送给相关的人，报警方式 比如  email sms weixin dingding  ;需要定义发送的信息 发送的方式
*/
// 接口  为了抽象行为

type Sender interface {
	Send(string) error
}

type AllSender interface {
	Send(string) error
	SendAll([]string) error
}

type SmsSender struct {
}

func (s *SmsSender) Send(msg string) error {
	return nil
}

type EmailSender struct {
	Addr     string
	Port     string
	User     string
	Password string
	To       string
}

// 方法
func (s *EmailSender) Send(msg string) error {
	fmt.Println("email sender:" + msg)
	return nil
}

// 方法
func (s *EmailSender) SendAll(msg []string) error {
	fmt.Println("email sender all:", msg)
	return nil
}

// 函数
// func test(sender *EmailSender) {
// 	sender.Send("yyyy")
// }

func test(sender *DingDingSender) {
	sender.Send("yyyy")
}

// 函数
func test2(sender Sender) {
	sender.Send("yyyy")
}

type DingDingSender struct {
}

func (s *DingDingSender) Send(msg string) error {
	fmt.Println("dingding sender:" + msg)
	return nil
}

func (s *DingDingSender) SendAll(msg []string) error {
	fmt.Println("dingding sender all:", msg)
	return nil
}

func main() {
	// var sender *EmailSender = new(EmailSender) //声明变量
	// sender.Send("xxxx")                        //email sender:xxxx #调用方法 ？？ 这个看前面的内容
	// test(sender)                               //email sender:yyyy

	// var sender *DingDingSender = new(DingDingSender) //声明变量
	// sender.Send("xxxx")                              //dingding sender:xxxx
	// test(sender)                                     //dingding sender:yyyy
	// //业务中，写好了一个EmailSender  如果后续要修改为 DingDingSender 则需要修改代码中的各个部分，这样不太好，能否抽象一个新的内容？

	// sender := new(EmailSender) // 声明
	// fmt.Println(sender.Addr)  // 为空值 #属性的调用方式
	// fmt.Printf("%T,%#v", sender, sender) *main.EmailSender,&main.EmailSender{Addr:"", Port:"", User:"", Password:"", To:""}

	// var sender Sender = new(EmailSender)
	// sender.Send("啦啦啦啦啦") // email sender:啦啦啦啦啦

	// var sender Sender = new(DingDingSender)
	// sender.Send("啦啦啦啦啦") // dingding sender:啦啦啦啦啦
	// 通过90 93 这两段的对比，发现定义了Sender接口后，我想修改不同的报警方式，在93行 初始化一个DingDingSender就好。

	// var sender Sender = new(EmailSender)
	// sender.Send("啦啦啦啦啦") // email sender:啦啦啦啦啦
	// // sender.SendAll([]string(""))  //sender 中没有SendAll  故会报错

	// var allsender AllSender = new(EmailSender)
	// allsender.Send("啦啦啦啦啦")   //email sender:啦啦啦啦啦
	// allsender.SendAll([]string{"aaaaaaaaaa"})  //email sender all: [aaaaaaaaaa] # 切片的写法
	// 通过97 101  这两段的对比发现，Sender 接口有两个方法（行动，这个方法是动词，不是数据类型方法），AllSender 有两个方法

	// sender.Send("xxx")
	// test2(sender)
	// sender, allsender
	// allsender = sender
	// sender = allsender
	// sender.Send("xxxxxx")
	// fmt.Println(allsender, sender)

}

//dingding结构体的指针  sender := new(DingDingSender)   // test2 在上面有一个赋值的过程

// 值对象和指针类型都可以赋值
// 接口不能访问对象的属性

//allsender = sender 不行
//sender = allsender 可以
// 区别是看里面的方法是否有

//接口变量赋值的目的是为了 抽象

// 断言  类型转回到结构体
