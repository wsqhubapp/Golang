package main

import "flag"

func main() {
	var (
		src string
		dst string
	)

	flag.StringVar()

	flag.Parse()

	// 检查
	/*
	 src
	 没有输入
	 文件路径不存在
	 目录

	 dst
	 没有输入
	 父目录不存在   =》 自动创建父目录（创建失败/创建成功）
	 文件存在
	 	目录

	*/
}
