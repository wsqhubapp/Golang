1 counter.go 计数 上一周的内容的扩展
2 作业 cpdir.go  作业的思路讲解
3 change 管道的补充
4 net.go  net包的学习
5 tcp     功能添加
5.1 tcpv1 基本功能
6 clock   tcp 并行
7 udpclock 
8 web      http 包实现了 类似于nginx的 web容器  !!!复制了文件
9 staticweb    路径  /static   相当 name/static/
10 webclient  用作爬虫   模拟浏览器




vs 中 深入的去看代码 如何实现 ？？？？

1. 复习
闭包陷进  传参来解决

为了解决共享数据的争抢问题 引入了锁和原子操作

延迟关闭锁 defer

var channel chan type   type是管道要传输数据的类型


多路复用

sync.Pool 空接口需要断言
go中没有泛型 （暂时）



不用关注历程的数量   历程是go来控制  不是操作系统来控制的，当历程超过百万级别的时候  可以考虑池的概念
2. 作业
作业中  第一个 第二个 抽象 就变成了 管道中 需要传递函数  函数处理复制也能处理字数统计
3. 新内容
    网络
        OSI七层参考模型
            应用 => 应用，表现(表示)，会话
            传输
            网络
            数据链路
            物理
        TCP/IP协议组 工业上进行验证
            操作系统的角度

            应用 HTTP, DNS, SMTP, RTMP, FTP, SSH
            传输 操作系统 TCP/UDP 源端口和目的端口  => 数据对应的程序（IP,端口）
            网络 IP, ICMP, ARP, RARP                 数据对应的电脑
            网络接口 网卡

网络分层 引到程序的分层   在第五天 讲到了user的程序

进程之间的通讯是通过网络实现

面向连接    打电话场景
非面向连接  短信场景

        socket 接口 操作系统提供给应用的网络开发接口

        telnet 工具

        TCP/UDP
        网络开发
            服务端/客户端

            服务端: 监听服务(装电话)
                    接收客户端连接 (等待电话)
                    交换数据
                    关闭客户端连接
                    关闭服务

            客户端:
                创建连接(拨打电话)
                交换数据
                关闭连接

            交换数据: 应用层数据格式(应用层协议)
                前5个字节 长度 99999
                read 5个字节 -> string -> int
                ctx make([]byte, length)

                写 => 前5个字节 => length

                []byte("3") => 00003
                []byte("00003")
                fmt.Printf("%5d")

                length = make([]byte, 5)

                append()

                abc

                字符串: 不包含换行
                    发送: 保证以换行结尾
                    接收: 带缓冲IO, ReadString("\n")

        net包
            工具

        并发：
            时间服务器，
            当client连接，给客户端响应一个服务器当前时间


        HTTP
        HTTP服务端&Client开发
            url => 处理逻辑
                    处理器/处理器函数
        RPC
        RPC服务端&Client开发

###关注点  vs code的使用问题  1 自动补齐  2 自动提示（自动提示接下来要输入的） 3鼠标放在某个词语上有显示这个怎么用