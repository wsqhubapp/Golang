// 讲解递归的问题
// 使用阶乘来做例子
package main

import "fmt"

func factV1(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	} else {
		rt := 1
		for i := 1; i <= n; i++ { // 关键思路停留在1*2*3*4...n 的形式上
			rt *= i
		}
		return rt
	}
}

/*
阶乘

n! = 1 * 2 * 3 * ... * n => (1* 2 * ...n-1)*n =>(n-1)!*n
0！= 1
递归的思路
递归是指函数直接或间接调用自己， 递归常用与解决分治问题，将大问题分解为相同的小问题进行解决，需要关注终止条件
或者说关键是 找到重复子问题
n < 0; n! 是定义错误
n =0  n！=1
n>0  n! = n* (n-1)!
fact(n) = n* fact(n-1)  // 从这行开始是伪代码的逻辑实现
递归  =》 结束条件
fact(3) = 3 * fact(2)
fact(2) = 2 * fact(1)
fact(1) = 1 * fact(0)
fact(0) = 1
*/
func fact(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(1))
}
