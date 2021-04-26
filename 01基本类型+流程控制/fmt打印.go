package main

import "fmt"

func main() {

	var a int = 123
	var b string = "abc"
	var c int64 = 123
	fmt.Println("println => 输出普通字符串或变量",a,b)
	fmt.Printf("pringf => 数值%d，字符串%s，变量c的类型是 %T\n",a,b,c)

}
