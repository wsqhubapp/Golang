20201008
数组的元素操作中没有 删除操作 因为数组是固定长度的
而切片的元素操作有删除操作  是因为切片是可变长度的

数组的切片操作 是和数组共享存储空间

需要再看看的
 数组的字面量
 数组的cap
 
 
切片的容量  cap-start


函数

定义
func 函数名（行参） 返回值{
 函数体
}

函数名  满足标识符规范

形参： 

返回值
 无返回值
 
调用



递归    解决分治问题



//递归伪代码
func rescursion(level, param1, param2, ...) {
    //中止条件
    if level > maxLevel {
     process_result
        return

    }
    //处理当前逻辑
    process(level, data...)
    //下探到下一层
    rescursion(level+1, p1, p2, ...)
    //清理当前层
}


找最近重复子问题

汉诺塔游戏