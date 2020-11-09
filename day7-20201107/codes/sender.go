package main


// 接口  为了抽象行为
//产生告警 发送给相关人

type EmailSender struct {

}

// 方法
func (s *EmailSender)


func test(sender * EmailSender){
	sender.Send("yyyy")
}

func main() {

}


//dingding结构体的指针  sender := new(DingDingSender)   // test2 在上面有一个赋值的过程

// 值对象和指针类型都可以赋值
// 接口不能访问对象的属性


//allsender = sender 不行
//sender = allsender 可以
// 区别是看里面的方法是否有

//接口变量赋值的目的是为了 抽象

// 断言  类型转回到结构体

