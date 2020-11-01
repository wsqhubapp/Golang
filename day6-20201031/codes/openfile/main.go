// 7 文件 追加
package main

import (
	"fmt"
	"time"
)

/*
O_RDONLY int = syscall.O_RDONLY // open the file read-only.
O_WRONLY int = syscall.O_WRONLY // open the file write-only.
O_RDWR   int = syscall.O_RDWR   // open the file read-write.
// The remaining values may be or'ed in to control behavior.
O_APPEND int = syscall.O_APPEND // append data to the file when writing.
O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O. 一般不使用
O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.（清空）

*/

func main() {
	os.ope

	if err != nil {
		return
	}

	defer file.Close()

	fmt.Fprintf(file, "%s\n", time.Now().Format("2006-12-06 15:44"))

}
