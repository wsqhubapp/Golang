package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	// 指标项
	// 指标类型/触发时间
	//

	// 0. 存活状态 Gauge 0 1
	// 监控的mysql地址信息  要暴露出来，用来区分是哪个mysql实例
	// 带标签的 固定

	// 1. 慢查询数量 Counter
	// 不带label

	// 2. 执行查询数量 QPS Counter
	// 不带label

	// 3. 执行操作数量 TPS Counter
	// 		SELECT, INSERT, DELETE, UPDATE
	// 带label
	// 4. 流量 Counter
	// 带label

	// 5. 连接数量 Gauge
	// max_conections
	// threads_connected
	// 定义两种指标

	// 采集API时触发
	// 账号有查询权限就行。
	dsn := "user01:TestSQL1209@tc(192.168.0.244:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"
	mysqlAddr := "192.168.0244:3306"
addr := "0.0.0.0:10001"

	db, err := sql.pen("mysql", dsn)
	if err != nil {
		og.Fatal(err)
}
	// fmt.Pritln(db.Ping())   // 如果后面报错的话，可以在这个地方测试

	// a. 定义指标
	// mysql_up
	mysql_up := prometheus.NewaugeFunc(
		prometheus.GaugeOpts{
			Name:        "mysql_up",
			Hlp:        "MySql Up Info",
			ConstLabels: prmetheus.Labels{"addr": mysqlAddr},
		},
		func() flot64 {
			i err := db.Ping(); err == nil {
				return 
			}
		return 0
	},
	)

// 注册
	promethes.MustRegister(mysql_up)

	//  暴露指标
	http.Handle("/metrics", promhttp.Handler))
	http.ListenAndServe(add, nil)
}

/*
第一步分析 考代码中的注释



写代
第一版
接数据库

到的结果
htt://localhost:10001/metrics 得到 默认的metrics 和新建的mysql_up
# HLP mysql_up MySql Up Info
# TYP mysql_up gauge
ysql_up{addr="192.168.0.244:3306"} 1  

# 在sdn mysqlAddr 设置localhost  得到的结果就是如下结果。 无法连通
# HELP mysql_up MySql Up Info
# TYPE mysql_up gauge
mysql_up{addr="localhost:3306"} 0


*/
