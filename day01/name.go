/*
标识符
规范：
1  必须满足 组成智能是非空的unicode的编码字符串，数字，下划线组成  // 中文也是可以的，比如我的名字，但是不建议用
2 必须满足  必须以unicode编码的字符串或下划线开头（不能以数字开头）
3 必须满足  不能与go的关键字冲突  比如  package func var 等25个关键字

建议:
1 ASCII 编码（a-z,A-Z） 数字  下划线
2 变量使用驼峰式  myName
3 与go 内置的标识符不要冲突（string）

说明：
go里面的标识符是区分大小写的   my My 是不同的变量

*/

/*
关键字
关键字用于特定的语法结构，go语言定义25关键字
声明， import package
实体声明和定义  chan const func interface map struct type var
流程控制  break case continue default defer else fallthrough for go goto if range return select switch
*/
