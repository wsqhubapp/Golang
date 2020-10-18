1 包
    包名： 标识符规范
    小写英文字母，尽量短小
    同一文件夹内所有go文件的包名必须一致
2 包管理
  gopath + vendor （历史）
  gomod（推荐 现在） 



  包名

        非main
            math => 库

升级不好弄 ，所以有 gomod


空（默认）   程序自己选择  
                  程序不在GOPATH 并且 当前目录有go.mod  就用gomod  否则使用GOPATH



go.mod 优势
1 版本写进文件
2 GOPROXY  
3 使用GOPATH 要不然放到固定的目录（GOPATH）， 否则需要配置新的GOPATH
3.1 go mod 代码可以随意放置在任何位置
4 go replace
    google.cn/aa --> github.com/aaa



初始化模块
go mod init 项目名称
    代码仓库的路径/项目名称
    一般命名方式  github.com/ismsilence/testgomod


    go 本质是使用git/svn 来组织第三方包  => 自动下载


    a 使用第三方包  // 下载到本地调用
    b 提供第三方包  //  go mod init  github.com/imsilence/testmath
    c 自己正常使用  //  go mod init 项目名称

    和外网不通的情况如何处理
    1 军工企业 
    自己内网，没有依赖任何第三方库
    依赖内部其他团队提供的库
    依赖外网的第三方库
        a. 编译环境  目标站点
                         acl
                         代理（proxy）  或者用镜像
                        把第三方包下载到  自己的gitlab/gitlee  // 要使用replace

go mod edit -replace=golang.org/x/net@latest=github.com/golang/net@latest
文件中实际的样子
replace golang.org/x/net latest => github.com/golang/net latest


导入
绝对导入
相对导入 只是在go path 支持  go mod 已经不支持了

点导入 # 尽量少用；当多个包都有Add 方法的时候，就不知道调用哪个包了
import (
    . "fmt"
)

别名导入 （解决多个相同包名的情况）
别名  实际路径 

下划线导入
导入不使用的包，和变量_ 作用一样； 写法和别名导入格式一样 
3 标准包
add 程序 可以接收多个int格式的字符串，至少输入两个  少于两个给提示
程序使用方法为 add 1 2 [3 4 5]
输入不是数字格式的字符串时，提示 程序使用方法为 add 1 2 [3 4 5]
计算所有数字总和 并输出结果

os
flag
time

    时间转换

hash  加密
base64
编码  []byte  string
解码  

go doc fmt
go doc fmt.Printf


4 第三方的包
5 单元测试

功能测试 testing 包
性能测试 也叫基准测试  



第三方库
godoc.org
go.dev

tablewriter

