// 方法的讲解 定义  初始化
package main

import (
	"fmt"
	"methods/models"
)

func main() {
	user := models.NewUser(1, "kk", 32)
	fmt.Println(user)   //{1 kk 32}
	models.AddAge(user) // 修改属性
	user.AddAge()
	fmt.Println(user)                 //{1 kk 32}
	fmt.Println(models.GetName(user)) //kk
	fmt.Println(user.GetName())       //kk

	puser := &user // 取引用 取地址
	// 调用函数如何调用
	(*puser).AddAge()               // 解引用 取值
	fmt.Println(*puser)             //{1 kk 32}
	fmt.Println((*puser).GetName()) //kk

	// GO的语法糖
	puser.AddAge()               // 自动解引用 (*puser).AddAge()
	fmt.Println(*puser)          //{1 kk 32}
	fmt.Println(puser.GetName()) //自动解引用 //(*puser).GetName()  //kk
}
