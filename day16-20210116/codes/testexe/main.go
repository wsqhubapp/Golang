package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	// "ls -la /"
	// dir .
	// cmd := exec.Command("ls", "-la", "/")
	fmt.Println(exec.LookPath("ping"))
	cmd := exec.Command("ping", "www.xxx.com", "-c", "1") // 如果是多个参数，或多个命令，写成 shell 来调用,这是因为 exec 里面的管道限制比较严格
	// output, err := cmd.Output()
	stdout, _ := cmd.StdoutPipe()

	// 异步处理
	cmd.Start()
	fmt.Println("started")

	io.Copy(os.Stdout, stdout)
	cmd.Wait()

	fmt.Println(cmd.ProcessState.ExitCode())
	os.Exit(100)

	// exex.Command(), Output()  cmd.ProcessState.ExitCode()
	// ->sh
}

// 第三方包  exec
/*
func lookpath

cmd  Output Run

stderrpipe
stderrpipe

stdoutpipe   管道 类似于 echo www.baidu.com |xargs ping


*/
