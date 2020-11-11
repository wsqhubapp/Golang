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
	// // 上面3行 这是直接访问指针（内存地址） 所以可以访问

	// combindVar := models.CombindStruct{}
	// fmt.Println(combindVar.PublicStruct.PublicAttrPu) //结果为空
	// fmt.Println(combindVar.PublicAttrPu)              //结果为空
	// //fmt.Println(combindVar.privateStruct.PublicAttrPr) //cannot refer to unexported field or method privateStruct
	// fmt.Println(combindVar.PublicAttrPr) //结果为空
	// // 嵌套的话，能访问到里层的大写属性

}
