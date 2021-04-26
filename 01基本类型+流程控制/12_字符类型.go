package main

import "fmt"

func main() {
	var ch byte // 什么字符类型
	ch = 97
	fmt.Printf("%c , %d\n",ch,ch)

	ch = 'a'
	fmt.Printf("%c , %d\n",ch,ch)

	fmt.Printf("%c\n",'A'+32)  // 大写转小写

}
