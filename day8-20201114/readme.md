课程顺序
tailf  目录是上周的作业讲解
1 goruntine.go  go例程的入门 
2 sleep.go      sleep的说明    让出调度
3 task.go
4 closer.go  
并发程序的通信方式
共享数据（同步）和管道（异步）
5 salary.go
salary_v1.go

6 channel.
7 bufferchannel.go
taskchannel.go
only.go  只读管道 只写管道
selectcase.go  多路复用  多个管道操作   只要有一个操作成功

timeout.go  历程的应用
time.go     time包讲解  time.After  多长时间后执行   time.Tick 间隔多长时间执行一次
timeout_v1.go
sync.go     sync 包讲解
runtime.go   runtime 包











具体内容
作业
第三方包的网站 go.dev  pkg.go.dev     go 的project 中也有这些介绍

tail -f
思路  有了思路后，有了伪代码  可以搜索 尝试解决问题
tail -n  就能用到bufferio





对于 顺序  并发 并行中的  并发 不理解
并发  cpu 分为多个时间片 
并行是并发的一种特殊情况
所以并发编程适用 现在所有的cpu环境（单核和多核）

并行编程  尽量让CPU用起来

操作系统层面  进程 线程
进程  资源分配的基本单位
线程  CPU调度的基本单位

go 并发编程参考
https://www.cnblogs.com/itbsl/p/12165215.html

例程通讯




应用   业务



作业：


