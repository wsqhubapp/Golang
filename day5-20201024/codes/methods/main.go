// 方法的讲解 定义  初始化
package main

import (
	"fmt"
	"methods/models"
)

func main() {
	user := models.NewUser(1, "kk", 32)
	fmt.Println(user)   //{1 kk 32}
	models.AddAge(user) // 修改属性 #这个调用里面的函数，相当于一个赋值的过程， 修改了里面的值，对外面的值是没有影响的。
	user.AddAge()
	fmt.Println(user)                 //{1 kk 32}
	fmt.Println(models.GetName(user)) //kk
	fmt.Println(user.GetName())       //kk

	puser := &user // 取引用 取地址或者说得到的是地址
	// 调用函数如何调用
	(*puser).AddAge()               // 解引用 取值或者说得到的是值
	fmt.Println(*puser)             //{1 kk 32}
	fmt.Println((*puser).GetName()) //kk

	// GO的语法糖  编译过程中，puser.AddAge() 会编译成为(*puser).AddAge()  #这种语法糖仅仅适用结构体调用方法的时候
	puser.AddAge()               // 自动解引用 (*puser).AddAge()
	fmt.Println(*puser)          //{1 kk 32}
	fmt.Println(puser.GetName()) //自动解引用 //(*puser).GetName()  //kk
}

// methods 和methodsv1 对比来看，能不能修改，取决于接收者  接收者是指针类型（引用类型） 可以修改
