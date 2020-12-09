package main

import (
	"fmt"
	"net/http"

	"user/config"
	"user/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	dsn := "user01:TestSQL1209@tcp(192.168.0.244:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"

	if err := config.InitDb(driverName, dsn); err != nil {
		fmt.Println(err)
		return
	}
	defer config.CloseDb()

	addr := ":8080"

	routers.Register()
	http.ListenAndServe(addr, nil)
}
