package main

import (
	"fmt"
	"methodsv1/models"
)

func main() {
	user := models.NewUser(1, "kk", 32)
	fmt.Println(user) //&{1 kk 32}
	// models.AddAge(user) // 修改属性
	user.AddAge()     // 结构的对象.方法名称()
	fmt.Println(user) //&{1 kk 33}
	// fmt.Println(models.GetName(user))
	fmt.Println(user.GetName()) //kk

	u2 := *user
	fmt.Printf("%T, %#v\n", u2, u2) //models.User, models.User{id:1, name:"kk", age:33}

	(&u2).AddAge()
	fmt.Println(u2)              //{1 kk 34}
	fmt.Println((&u2).GetName()) //kk
	// Go语法糖
	u2.AddAge()               // (&u2).AddAge()
	fmt.Println(u2)           ////{1 kk 35}
	fmt.Println(u2.GetName()) // (&u2).GetName() //kk
}
