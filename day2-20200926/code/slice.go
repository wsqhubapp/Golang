// 切片
package main

import "fmt"

func main() {
	// 定义切片
	var names []string         // 这个是nil； nil和空切片并不是一样的
	fmt.Printf("%T\n", names)  // []string
	fmt.Printf("%#v\n", names) // []string(nil)
	fmt.Printf("%q\n", names)  // []
	fmt.Println(names)         // []

	names = []string{"abc", "sen"}
	fmt.Printf("%T\n", names)
	fmt.Printf("%#v\n", names)

	// 字面量
	// []type{}  这个是空切片
	// []

	names = []string{1: "abc", 10: "sen"}
	fmt.Printf("%T\n", names) // []string
	fmt.Printf("%q\n", names) // ["" "abc" "" "" "" "" "" "" "" "" "sen"]

	// 访问 修改元素
	// 索引
	// *  超过范围的话
	fmt.Println(names[1])
	fmt.Println(names[10])

	// 长度
	fmt.Println(len(names)) // 11

	// 遍历
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	for  v1 := range names {
		fmt.Println(v1 names[v1])
	}

	for i, v := range names {
		fmt.Println(i, v)
	}


	// 
	names =  append

	// 删除元素  切片操作
	names[start:end] names 中从start 开始到end -1 所有元素组成的切片 不包含end
	




}
