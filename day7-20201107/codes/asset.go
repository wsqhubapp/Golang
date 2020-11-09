// 类型的 断言 和查询
// 断言是一个抽象到具体的问题   接口 结构体的指针之间的操作
package main


//duanyan
obj, ok := sender.(*EmailSender)
fmt.Println(obj, ok)



//查询
//swich case  一个一个去匹配 能匹配上就匹配第一个   从上往下匹配
//我以为是 sender - allsender - email 的顺序
//但实际上是switch语法本身决定了顺序
//sender.(type) 是一个固定语法

//接口是多态的一个思路
