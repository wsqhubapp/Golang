/*
值类型和引用类型的差异在于赋值同类型新变量后，对新变量进行修改是否能够影响原来的变量，如果不能影响，则为值类型，如果影响则为引用类型

*/
package main

import "fmt"

func main() {
	name, age, heigh, isBoy := "silence", 30, 168, false
	pointer := new(int)           //定义指针类型
	scores := [...]int{1, 2, 3}   // 定义数组类型
	names := make([]string, 1, 3) // 定义切片类型
	user := make(map[int]string)  //定义映射类型

	name2, age2, heigh2, isBoy2, pointer2, scores2, names2, user2 := name, age, heigh, isBoy, pointer, scores, names, user

	name2 = "kk"
	age2 = 31
	heigh2 = 170
	isBoy2 = true
	scores2[0] = 0
	pointer2 = &age
	names2[0] = "kk"
	user2[1] = "kk"

	fmt.Println(name, age, heigh, isBoy, scores, pointer, names, user)
	fmt.Println(name2, age2, heigh2, isBoy2, scores2, pointer2, names2, user2)
	/*
		silence 30 168 false [1 2 3] 0xc0000120b8 [kk] map[1:kk]
		kk      31 170 true  [0 2 3] 0xc0000120b0 [kk] map[1:kk]
		有上面对比，可知道 后面的切片和映射影响了原来的值
		值类型： 数值、布尔、字符串、指针、数组、结构体等
		引用类型： 切片 映射 接口等
	*/
	fmt.Printf("%p,%p,%p,%p,%p,%p,%p,%p\n", &name, &age, &heigh, &isBoy, &scores, &pointer, &names, &user)
	fmt.Printf("%p,%p,%p,%p,%p,%p,%p,%p\n", &name2, &age2, &heigh2, &isBoy2, &scores2, &pointer2, &names2, &user2)
	/*
		针对值类型和引用类型在赋值后新旧变量的地址并不相同，只是引用类型在底层共享数据结构（其中包含指针类型元素）
		0xc00003c1f0,0xc0000120b0,0xc0000120b8,0xc0000120c0,0xc000010440,0xc000006028,0xc0000044a0,0xc000006030
		0xc00003c200,0xc0000120d0,0xc0000120d8,0xc0000120e0,0xc000010460,0xc000006038,0xc0000044c0,0xc000006040

		针对值类型可以借助指针修改原值
	*/

}
