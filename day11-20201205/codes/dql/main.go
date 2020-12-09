package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	// "user:password@protocol(host:port)/dbname?charset=utf8mb4&loc=Local&parseTime=true"        //data store name 数据库连接信息，使用协议，用户&密码，数据库，连接参数

	dsn := "user01:TestSQL1209@tcp(192.168.0.244:3306)/user?charset=utf8mb4&loc=Local&parseTime=true" //data store name 数据库连接信息，使用协议，用户&密码，数据库，连接参数

	db, err := sql.Open(driverName, dsn) // 建立连接
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close() // 连接关闭

	if err := db.Ping(); err != nil { // 连通测试  如果连通出错的话 会报错
		fmt.Println(err)
		return
	}

	name := "%kk%" // 数据库的预处理方式来防止sql注入
	sql := `
		select id, name, password, sex, birthday, addr, tel
		from user
		where name like ?
		order by ? desc
		limit ? offset ?
	`
	fmt.Println(sql)
	// 操作  查询的操作
	rows, err := db.Query(sql, name, "birthday", 3, 3) // 数据库的预处理方式  每个参数代替 上面的问号
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id       int64
			name     string
			password string
			sex      bool
			birthday *time.Time
			addr     string
			tel      string
		)
		err := rows.Scan(&id, &name, &password, &sex, &birthday, &addr, &tel)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(id, name, password, sex, birthday, addr, tel)
		}
	}

	var id int64
	err = db.QueryRow("select id from user").Scan(&id) // 查询一行

	fmt.Println(err, id)
}

/*
密码中可以有@  程序中只认最后一个@
time.Time 是一个非空的 字段


数据库的预处理方式
关键字（表明、字段名、select等自带的语句 ）不能替代，其他的可以替代
*/
