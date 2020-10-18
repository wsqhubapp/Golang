package main

import (
	"fmt"
	"os"
	"strconv"
)

var users = []map[string]string{}

/*
最终得到这样样式
[map[address:suzhoushi id:1 name:wsq phone:151] map[address:wuxishi id:2 name:w1 phone:139]]
*/

//将相同的逻辑提取为函数, 提高代码简洁性

func commonInput() (string, string, string) {
	var name, phone, address string
	fmt.Printf("请输入用户名字: ")
	fmt.Scan(&name)
	fmt.Printf("请输入用户联系方式: ")
	fmt.Scan(&phone)
	fmt.Printf("请输入用户通信地址: ")
	fmt.Scan(&address)

	return name, phone, address
}

/*
	增
	从命令行分别输入名称、联系方式、通信地址
	生成ID => 查找users中最大的id+1（无元素id=1） 放入到users
*/
func add() {
	name, phone, address := commonInput()

	/*
		匿名函数
		slice为空ID=0, slice不为空，新增时slice中最大的ID+1
		切记不可len(slice)+1, 这样会有个bug，当删除一个元素时，再新增会导致相同ID的元素有两个
	*/
	maxId := func(users []map[string]string) string {
		if len(users) == 0 {
			return "1"
		} else {
			for i := 0; i < len(users)-1; i++ {
				if users[i]["id"] < users[i+1]["id"] {
					users[i], users[i+1] = users[i+1], users[i]
				}
			}
		}
		max, _ := strconv.Atoi(users[0]["id"])
		return strconv.Itoa(max + 1)
	}(users)
	users = append(users, map[string]string{
		"id":      maxId,
		"name":    name,
		"phone":   phone,
		"address": address,
	})

	fmt.Println(users)
}

/*
	删
	删除 del函数 从命令行中输入要删除的用户ID 验证ID是否存在，如果存在，打印需要删除用户信息 并让用户输入y/n确认是否删除 输入y删除用户信息
	fmt.Println(users)
*/
func del() {
	var id string
	fmt.Printf("请输入要删除的用户Id: ")
	fmt.Scan(&id)

	for userIndex, user := range users {
		value, ok := user["id"]
		if ok && value == id {
			fmt.Println(user)
			var isDel string
			fmt.Printf("是否删除 y/n: ")
			fmt.Scan(&isDel)
			if isDel == "y" {
				users = append(users[:userIndex], users[userIndex+1:]...)
				fmt.Println(users)
			}
		}
	}
}

/*
	改
	修改 modify函数 从命令行输入要修改的用户ID 验证ID是否存在，如果存在，打印需要修改用户信息 并让用户输入y/n确认是否修改 输入y修改用户信息，继续让从命令行分别输入 用户名，联系方式，地址 进行更新
	fmt.Println(users)
*/
func modify() {
	var id string
	fmt.Printf("请输入要修改的用户Id: ")
	fmt.Scan(&id)

	for _, user := range users {
		value, ok := user["id"]
		if ok && value == id {
			fmt.Println(user)
			var isModify string
			fmt.Printf("是否确认修改 y/n: ")
			fmt.Scan(&isModify)
			if isModify == "y" {
				name, phone, address := commonInput()
				user["name"] = name
				user["phone"] = phone
				user["address"] = address
				fmt.Println(users)
			}
		}
	}
}

/*
	查
	查找 query函数 从命令行输入要查询的字符串 遍历比较用户的名称，地址，联系方式，包含要查找的字符串就进行输出
	Id：xxx 名字：xxx
*/
func query() {
	var queryStr string
	fmt.Printf("请输入要查询的字符串: ")
	fmt.Scan(&queryStr)
	for _, user := range users {
		for userKey, userInfo := range user {
			if userInfo == queryStr {
				fmt.Printf("Id: %s, %s: %s \n", user["id"], userKey, userInfo)
			}
		}
	}
}

// display  显示现在的用户信息
func display() {
	fmt.Println(users)
}

func main() {

	fmt.Printf(`
	1. 添加用户 ==> add
	2. 删除用户 ==> del
	3. 修改用户 ==> modify
	4. 查询用户 ==> query
	5. 显示用户 ==> display
	6. 退出 ==> quit
	`)

	for {
		var choiceOne string
		fmt.Printf("请选择你要做的操作: ")
		fmt.Scan(&choiceOne)
		switch choiceOne {
		case "add":
			add()
		case "query":
			query()
		case "del":
			del()
		case "modify":
			modify()
		case "display":
			display()
		case "quit":
			os.Exit(0)
		}
	}
}
