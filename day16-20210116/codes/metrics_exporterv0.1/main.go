package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	addr := ":9090"
	// 定义指标
	totalV1 := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace:   "",
			Subsystem:   "",
			Name:        "test_total_v1",
			Help:        "Test Total V1 counter",
			ConstLabels: map[string]string{"name": "v1"},
		},
	)

	// 注册指标
	prometheus.MustRegister(totalV1)
	// 修改指标值
	totalV1.Add(10)
	// 暴露httpapi
	http.Handle("/metrics", promhttp.Handler())

	// 启动web服务
	http.ListenAndServe(addr, nil)

}

/*
启动 go run main.go
在浏览器中 打开http://localhost:9090/metrics   能看到下面的内容
# HELP test_total_v1 Test Total V1 counter
# TYPE test_total_v1 counter
test_total_v1{name="v1"} 10

*/
