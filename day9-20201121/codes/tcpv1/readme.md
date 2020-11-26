1 实现了基本的数据发送
2 服务器 建立了服务的5个过程，客户端实现了建立服务的3个过程
服务端: 监听服务(装电话)    listener     net.Listen
        接收客户端连接 (等待电话)  conn listener.Accept() 
        交换数据        read(conn)  write(conn)
        关闭客户端连接  conn.close()
        关闭服务 listener.close()

客户端:
    创建连接(拨打电话)    conn net.Dial()
    交换数据   read()  write()
    关闭连接  conn.close

3 连接成功后 多次的交换数据  在交换数据的逻辑前面  添加for

