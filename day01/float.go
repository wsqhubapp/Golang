package main

import "fmt"

func main() {
	var height float32 = 1.68
	var heightType = 1.68 // 科学计数法

	fmt.Printf("%T, %#v %f\n", height, height, height)
	fmt.Printf("%T, %#v %f\n", heightType, heightType, heightType)

	// 操作
	// 算术运算 + - * / ++ --  // 没有取余的运算
	var (
		f1 = 1.2
		f2 = 2.36
	)
	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)
	fmt.Println(f1 / f2)
	/* 结果
	3.5599999999999996
	-1.16
	2.832
	0.5084745762711864

	*/

	// 关系运算 > < >= <=   //没有不等于或等于的运算。
	fmt.Println(f1 > f2)
	fmt.Println(f1 < f2)
	fmt.Println(f1 >= f2)
	fmt.Println(f1 <= f2)

	// 没有位运算

	// 赋值运算 =  += -= /=  *=
	f1 = 1.32
	f1 += f2 //实际表示 f1 = f1 + f2

	fmt.Println(f1)

}
