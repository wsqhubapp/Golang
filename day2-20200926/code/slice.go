// 切片
package main

import "fmt"

func main() {
	// 定义切片
	var names []string // 这个是nil； nil和空切片并不是一样的
	// fmt.Printf("%T\n", names)  // []string
	// fmt.Printf("%#v\n", names) // []string(nil)
	// fmt.Printf("%q\n", names)  // []
	// fmt.Println(names)         // []

	names = []string{"abc", "sen"}
	// fmt.Printf("%T\n", names)
	// fmt.Printf("%#v\n", names)

	// 字面量
	// []type{}  这个是空切片
	// []type{v1, v2, v3, ... vn}  值的方式
	// []type{i1:v1, i2:v2 ... in:vn}

	names = []string{1: "abc", 10: "sen"}
	// fmt.Printf("%T\n", names) // []string
	// fmt.Printf("%q\n", names) // ["" "abc" "" "" "" "" "" "" "" "" "sen"]

	// 访问 修改元素
	// 索引
	// *  超过范围的话
	// fmt.Println(names[1])  // abc
	// fmt.Println(names[10]) // sen

	// 长度
	// fmt.Println(len(names)) // 11

	// 遍历  下面3种不同的遍历方式，得到结果一样
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(i, names[i])  // 索引i 和value值
	// }

	// for v1 := range names {
	// 	fmt.Println(v1, names[v1])  // v1 得到的是索引值 一个参数的事
	// }

	// for i, v := range names {  //最直接的写法
	// 	fmt.Println(i, v)
	// }

	// 添加元素
	// 切片的末尾
	names = append(names, "liuran")
	// fmt.Printf("%q, %#v\n", names, len(names))  //["" "abc" "" "" "" "" "" "" "" "" "sen" "liuran"], 12

	// 删除元素  通过切片操作实现
	//names[start:end] names 中从start 开始到end -1 所有元素组成的切片 不包含end
	// names[1:10]
	//  1   2                              11    12
	// ["" "abc" "" "" "" "" "" "" "" "" "sen" "liuran"]
	// fmt.Printf("%q\n", names[1:10]) //["abc" "" "" "" "" "" "" "" ""]  //这样就去掉了后面2个
	// names = names[1:len(names)]
	// fmt.Printf("%q\n", names) //["abc" "" "" "" "" "" "" "" "" "sen" "liuran"] 没有去掉

	// names = names[0 : len(names)-1] //
	// fmt.Printf("%q\n", names)       //["abc" "" "" "" "" "" "" "" "" "sen"]  这样就去掉了后面一个

	//删除中间的元素
	//目标是删除索引为4的元素  用copy函数实现  copy(dst, src) src -> dst
	nums := []int{0, 1, 2, 3, 4, 5}

	nums2 := []int{10, 11, 12, 13, 14, 15, 16}

	copy(nums, nums2)

	fmt.Println(nums, nums2) //[10 11 12 13 14 15] [10 11 12 13 14 15 16]

	copy(nums[3:len(nums)], nums[4:len(nums)])
	// [13 14 15]   [14 15]

	fmt.Println(nums[0:len(nums)])
	fmt.Println(nums[0 : len(nums)-1])

}
