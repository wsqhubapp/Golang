1 复习
2 作业处理
    修改自己的密码
    1 打开修改密码的页面，让用户输入密码信息
        mvc
        router url =》 controller
        controller :
            GET
                ctx
                TPl


    2 输入密码信息，提交并修改
        旧密码
        新密码
        新密码确认

3 新内容

知识点：
    控制器：
        响应格式
        URL构建
        Flash 消息
        错误处理
        日志处理
        过滤器
    表单：
        验证
    模型：
        关系
    模板：
        函数使用
        Layout&LayoutSections
    安全：
        xsrf



        OAUTH
SSO
IAM
3A
4A
5A
都是用户管理系统


部署
    构建
        编译二进制程序
            set GOOS = linux
            go build
        
        压缩包 非go的都需要部署

        部署
            设置执行权限
            更改配置
            启动

        问题：
            1 命令行启动
                nohup
            2 systemd
                supervisord
            3 go 即处理webapi 又处理静态资源
            nginx + go 
            4 nginx 负载多个go进程
            
                redis

