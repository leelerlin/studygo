package main

import "fmt"

func main() {
	var a bool  // 默认是false（零值）
	a =  true
	fmt.Println("a = ", a)

	// 2.自动推导
	var b = false
	fmt.Println("b =", b)

	c := false
	fmt.Println("c = ",c)
}
