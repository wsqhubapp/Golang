// 指针类型 point  老师需要缩写为p
package main

import "fmt"

func main() {
	// 指针： 每个变量在内存中都有对应存储位置（内存地址），可以通过&运算符获取。指针是用来存储变量地址的变量
	// 值类型
	// 赋值 原有的数据复制一份给新的变量
	// 两个变量之间没有任何联系
	// 对nums进行任何修改都不会影响nums2
	// 对nums2进行任何修改也不会影响nums

	// 有联系 => 引用类型(内部使用使用了指针/地址)

	// 数组值类型
	nums := [5]int{1, 2, 3, 4, 5}
	nums2 := nums //声明

	fmt.Println(nums, nums2) //[1 2 3 4 5] [1 2 3 4 5]

	nums2[0] = 100           //赋值
	fmt.Println(nums, nums2) //[1 2 3 4 5] [100 2 3 4 5]  两个不同的数组，修改了其中一个，对另外一个是没有影响的
	// int, float, bool, string 前面几个说到的是值类型

	// 声明
	// var name *type

	var age = 1
	var agePointer *int = nil // 定义int类型的指针
	// 初始化 方法一  使用&运算符+变量初始化：&运算获取变量的存储位置（内存地址）来初始化指针变量
	// 初始化 方法二 使用 new 函数初始化：new 函数根据数据类型申请内存空间并使用零值填充，并返回申请空间地址
	// 操作 可通过*运算符+指针变量名来访问和修改对应存储位置的值

	agePointer = &age // 取出age的地址赋值给agePointer => 取引用

	fmt.Println(agePointer, age) //0xc0000120a8 1

	fmt.Println(*agePointer) //1  #获取agePointer内存地址中对应的存储的值 => 解引用

	*agePointer = 33              //赋值 # 这也是上面说到的操作的一种方式
	fmt.Println(age, *agePointer) //33 33 #age 和agepoint 使用的是同一个内存地址，所以两个值都是33

	var numsPoint *[5]int = &nums
	fmt.Printf("%T, %#v\n", numsPoint, numsPoint) //*[5]int, &[5]int{1, 2, 3, 4, 5}
	numsPoint[0] = 100                            //赋值
	fmt.Println(nums, numsPoint)                  //[100 2 3 4 5] &[100 2 3 4 5]

}
