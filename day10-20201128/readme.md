课程顺序
1 tcprpc  构建在tcp 协议上的RPC  
2 httprpc 
3 jsonrpc
4 webrequest   http 请求内容
5 webparams  http 提交参数



听老师讲课的心得
1） 要有思路    第一步 干嘛  第二步 干嘛
2） 代码书写  从已经知道到不知道的位置 比如 说 学一个新的包   是 有 io.getxx  ;我现在想打印出来， 然后在这个代码前面写适合的打印语句；而不是先写打印语句  后写实际的内容  


1. 复习：
    Read
        交换数据 只有一行且内容无\n \n
        如果用\n来结束流读取, 内容中不能有\n
2. 作业：
3. 课程内容：
    1. rpc
        rpc       底层数据编码使用gob
            http
            tcp
        jsonrpc   构建于tcp 协议之上，底层数据编码使用  json 可以跨越语言

        Register(&Calc)

    总结： 通用逻辑，具体的实现语句不同，底层用不同的方式编码 不需要我们具体去写，是一个Go语言包装后的结果


    2. http
        请求：
            数据接收&处理
        响应：
            text/template
            html/template
            模板技术
            fmt.Printf("xxx %s xxx", xxx, )
            字符串模板 + 数据 => 组装 字符串进行输出

            模板+数据 =>引擎 => 组装 字符串

    3. log

    机动:
    4. goquery
    5. regexp

