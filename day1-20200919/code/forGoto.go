package main

import "fmt"

func main() {
	// 1 + 2 +3 .. 10
	sum := 0
	idx := 1

START:
	if idx > 10 {
		goto END
	}

	sum += idx
	idx++
	goto START

END:
	fmt.Println(sum)
}

/*可以通过goto语句任意跳转到当前函数指定的label位置*/
