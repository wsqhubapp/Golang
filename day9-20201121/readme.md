1 counter.go 计数 上一周的内容的扩展
2 作业 cpdir.go  作业的思路讲解
3 change 管道的补充
4 net.go  net包的学习
5 tcp
6 clock  并行
7 udpclock 
8 web      http 包实现了 类似于nginx的 web容器
9 staticweb    路径  /static   相当 name/static/
10 webclient  用作爬虫   模拟浏览器






复习
闭包陷进  传参来解决


为了解决共享数据的争抢问题 引入了锁和原子操作

延迟关闭锁 defer

var channel chan type   type是管道要传输数据的类型


多路复用

sync.Pool 空接口需要断言
go中没有泛型 （暂时）



不用关注历程的数量   历程是go来控制  不是操作系统来控制的，当历程超过百万级别的时候  可以考虑池的概念
作业
1

作业中  第一个 第二个 抽象 就变成了 管道中 需要传递函数  函数处理复制也能处理字数统计


新的内容
网络分层 引到程序的分层   在第五天 讲到了user的程序

进程之间的通讯是通过网络实现

面向连接    打电话场景
非面向连接  短信场景

vs 中 深入的去看代码 如何实现 ？？？？
