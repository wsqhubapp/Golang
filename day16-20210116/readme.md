1. JS 课程自学
    javascript ES5
    jQuery
    jQuery插件
        Bootstrap JS
        datatables
        sweetalert
        notification

上面是js的内容

2. prometheus
    熟练使用 ps:1
    简单使用 ps:2
    不知道/没用过 ps:0
    老师还是讲了基础知识

1. 复习
2. 作业
3. 新内容
    prometheus
        搭建
        基本原理
        示例
        PromQL
        alertmanager
    exporter => mysql_exporter
    告警管理
    Agent => Prometheus配置应用

// json/ini

yaml 配置文件

注释 # 开头
key: value # key名字value值  标准格式

key:字符串
value: 类型
    数值类型 1, 1.1
    字符串类型 kk "kk"   #字符串不一定要用双引号标识； 但是比如 字符串true，就要用"true"
    布尔类型 true false
    列表类型 [1, 2, 3]
    key:
    - 1
    - 2
    - 3
    字典类型
    {key:value}
    key:
        subkey: value
        subkey: xxx

使用缩进表示层级关系，    2个空格（建议值）/4个空格  不允许使用tab键；缩进的空格数目不重要，只要相同层级的元素左侧对齐即可。
参考：https://cnblogs.com/wxmdevelop/p/7341292.html

webhook
    hook: 钩子
    在一个程序里面 执行流程 过程中通过调用外部接口 改变流程/调用外部接口进行通知

    webhook 对应接口是WEB API

systemd.service
    /usr/lib/systemd/system


prometheus：
    内置指标
        up 采集目标是否正常

如何监听操作系统信号
SIGHUP
代码 testsignal
执行系统命令如何执行  exec  代码testcmd

go中执行系统命令， 调用os的模块
进程类似的  gopsutil
参考：https://www.cnblogs.com/Dr-wei/p/11742378.html
https://github.com/shirou/gopsutil


下午内容
0s-1s之内 响应10次
1s-2s之内 响应10次
20 0.50 = 10


建议学习 图解HTTP 这本书

prometheus 安全认证  # prometheus 认证  https://prometheus.io/docs/prometheus/latest/configuration/https/
basic auth
    认证失败:
        401
        Www-Authenticate: Basic
        autherization: Basic XXXXXX

        base64(user:password)  # 这个到网上 能反向解析      base64 编码解码的过程

    nginx
    apache

    web basic


TOKEN: jwt token

autherization: Bearer TOKEN



go http web文件服务器




1. expoter

http://ip:9100/metrics

http
/metrics

ip:9100

配置文件 具体参数  需要去学习  总结

规则：
    记录
        使用当前采样指标 生成新的采样指标


    告警

cpu=0 instance=localhost:9100
cpu=1 instance=localhost:9100
cpu=0 instance=1.1.1.1:9100
cpu=0 instance=1.1.1.2:9100
cpu=1 instance=1.1.1.2:9100
cpu=2 instance=1.1.1.2:9100

count(*) group by instance

count(node_cpu_seconds_total{mode="idle"}) by (instance)  #查看有几个CPU的核心


node_cpu_seconds_total[5m]
node_cpu_seconds_total{mode="idle"}


时间区间：
    5m
    5s
    5h
    1d
    1w
    1y

一小时之前up


count(metric{label}) by ()


sum(1-irate(node_cpu_seconds_total{mode="idle"}[5m])) by (instance) * 100
(1 - node_memory_MemAvailable_bytes/node_memory_MemTotal_bytes) * 100    #内存的使用率









