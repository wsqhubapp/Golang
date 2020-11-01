// 9 文件读写
// EOF  end of file   这里不支持插入

package main

import "os"

func main() {
	os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)

}
