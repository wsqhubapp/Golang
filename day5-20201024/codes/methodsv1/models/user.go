package models

type User struct {
	id   int
	name string
	age  int
}

// 函数
func NewUser(id int, name string, age int) *User {
	return &User{id, name, age}
}

// 方法 user结构体指针接收者方法
func (user *User) GetName() string {
	return user.name
}

// 方法
func (user *User) AddAge() {
	user.age += 1
}

// 函数
func AddAge(user *User) {
	user.age += 1
}

//函数
func GetName(user *User) string {
	return user.name
}
