package main

import "fmt"

// 函数也是一种数据类型，通过type给一个函数类型起名
type FuncType2 func(int,int) int

func main() {
	ret := add(1,2)
	fmt.Println(ret)

	var ftest FuncType2
	ftest = add
	ret2 := ftest(3,4)
	fmt.Println(ret2)
}

func add(a,b int) int {
	return a+b
}

func minus(a,b int) int {
	return a-b
}

