/*
1 md5 sha1  sha256 sha512  的作用，用法类似，结果是长度不同而已
2 能实现多个字符拼接的效果
3 注意两种不同write的写法
*/

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	txt := "i am kk"
	md5Value := fmt.Sprintf("%x", md5.Sum([]byte(txt)))
	fmt.Println(md5Value)
	sha1Value := fmt.Sprintf("%x", sha1.Sum([]byte(txt)))
	fmt.Println(sha1Value)
	sha256Value := fmt.Sprintf("%x", sha256.Sum256([]byte(txt)))
	fmt.Println(sha256Value)
	sha512Value := fmt.Sprintf("%x", sha512.Sum512([]byte(txt)))
	fmt.Println(sha512Value)

	md5hasher := md5.New()
	md5hasher.Write([]byte("i am "))
	md5hasher.Write([]byte("kk"))

	fmt.Printf("%x\n", md5hasher.Sum(nil))

	// 计算字节切片md5散列值（md5.Sum()）
	fmt.Printf("%x\n", md5.Sum([]byte("Hi kk")))
	// 创建md5 hash对象（md5.New()）
	hash := md5.New()
	// 批量写入字符串计算散列值（ 注意空格和标点符号，空格写在第一个的后面和第二个的前面都行）
	io.WriteString(hash, "Hi ")
	io.WriteString(hash, "kk")

	fmt.Printf("%x\n", hash.Sum(nil))
}
