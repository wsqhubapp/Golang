// 结构体嵌套的话，可见性的问题
package main

func main() {
	// 	// fmt.Printf("%T\n", models.PublicStruct{}) //models.PublicStruct

	// 	publicVar := models.PublicStruct{}
	// 	fmt.Printf("%#v\n", publicVar)      //models.PublicStruct{privateAttrPu:"", PublicAttrPu:""}
	// 	fmt.Println(publicVar.PublicAttrPu) //结果为空
	// 	// fmt.Println(publicVar.privateAttr)  //privateAttr undefined
	// 	// fmt.Println(models.privateStruct{}) //cannot refer to unexported name models.privateStruct
	// //结构体大写的，属性名大写的能访问；结构体大写的，属性名小写的不能能访问； 结构体小写的，属性名大小写的都不能访问

	// privateVar := models.NewPrivateStruct()
	// fmt.Printf("%#v\n", privateVar)      //&models.privateStruct{privateAttrPr:"", PublicAttrPr:""}
	// fmt.Println(privateVar.PublicAttrPr) //结果为空
	// // fmt.Println(privateVar.privateAttr)  //privateVar.privateAttr undefined
	// // 上面3行 这是一个函数的调用。这里这么写，主要是想说明，7开头的一段，外面能访问，不好控制变量， 14开头的一段是函数的调用，可以做控制
	// // 这是从业务角度的一个考虑和练习

	// combindVar := models.CombindStruct{}
	// fmt.Println(combindVar.PublicStruct.PublicAttrPu) //结果为空
	// fmt.Println(combindVar.PublicAttrPu)              //结果为空
	// //fmt.Println(combindVar.privateStruct.PublicAttrPr) //cannot refer to unexported field or method privateStruct
	// fmt.Println(combindVar.PublicAttrPr) //结果为空
	// // 嵌套的话，能访问到里层的大写属性  这是因为中间省略了，而程序看到的都是大写的，可以访问

}
