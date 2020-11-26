package main

import (
	"log"
	"os"
)

// // 自定义日志格式
// func init() {
// 	log.SetPrefix("[usercenter]")  // 添加前缀
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)  // 时间 文件 等的设置  有哪些可以设置查看log 的const
// }

// 定义日志写到文件
func init() {
	file := "./" + "message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile) // 将文件设置为log 输出文件
	log.SetPrefix("[tool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return
}

func main() {
	// log.Println("我是一个兵")
	// log.Printf("我是一个%s \n", "兵")
	log.Println("Hello go")
}

/*
也有  log.fatal  等错误日志
*/
