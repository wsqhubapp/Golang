package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	driverName := "mysql"
	dsn := "user01:TestSQL1209@tcp(192.168.0.244:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"
	db, err := sql.Open(driverName, dsn)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	// 插入
	name, password, birthday := "kk2", "xxxx", "1999-11-11"
	sql := `
	INSERT INTO user(name, password, birthday) VALUES(?, ?, ?);
	`
	result, err := db.Exec(sql, name, password, birthday)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId()) // 插入的是哪一行
		fmt.Println(result.RowsAffected()) // 生效的有几行
	}

	// 更新
	sql = `
	UPDATE user
	SET birthday=now()
	WHERE id=?
	`
	result, err = db.Exec(sql, 1)
	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	// 删除
	sql = `
	DELETE from user where name=?
	`

	result, err = db.Exec(sql, name) // 这个name 在上面有定义 kk2
	fmt.Println(err)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
