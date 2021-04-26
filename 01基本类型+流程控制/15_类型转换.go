package main

import "fmt"

func main() {
	var flag bool
	fmt.Printf("flag = %t\n",flag)

	flag = true
	fmt.Printf("flag = %t\n",flag)

	//bool 类型不能转换为int
	//fmt.Printf("flag = %d \n",int(flag))

	// 整形也不能转换成bool
	//flag = bool(1)

	var ch byte = 'a' // 字符类型本质是整型
	var t int
	t = int(ch)
	fmt.Println("t=",t)

}
