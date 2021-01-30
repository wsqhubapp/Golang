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

	totalV2 := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "test",
			Subsystem:   "total",
			Name:        "V2",
			Help:        "Test Total V2 Counter",
			ConstLabels: prometheus.Labels{"name": "v2"},
		},
		[]string{"path"},
	)

	// 注册指标
	prometheus.MustRegister(totalV1)
	prometheus.MustRegister(totalV2)

	// 修改指标值
	totalV1.Add(10)
	totalV2.WithLabelValues("/root").Inc()
	totalV2.WithLabelValues("/login").Add(5)
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

# HELP test_total_V2 Test Total V2 Counter
# TYPE test_total_V2 counter
test_total_V2{name="v2",path="/login"} 5
test_total_V2{name="v2",path="/root"} 1

*/
