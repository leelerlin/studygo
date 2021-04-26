package main

import "fmt"

func main() {
	// if
	a,b,c := 1,2,3
	if a == 1 {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}

	// switch
	switch a {
	case 1:
		fmt.Println('1')
		// break // go语言保留了break关键字，默认存在，可以省略
		// fallthrough // 不跳出，继续执行后面的逻辑
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

	switch num:=2;num {
	case 1:
		fmt.Println('1')
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

	score := 80
	switch {
	case score > 90:
		fmt.Println("> 90")
	case score < 90:
		fmt.Println("<90")
	}

	//for  go语言循环只有for
	sum := 0
	for i,j:=0,1;i <=100 ; i++ {
		sum = sum + i
		sum = sum + j
	}
	fmt.Println(sum);

	// for 死循环
	//for {
	//	println("abc");
	//}

	// range
	str := "abc"
	for i,data := range str { // range 默认返回两个参数，第一个是参数下标，第二个返回下标值
		fmt.Printf("str[%d]=%c\n",i,data)
	}
	for i := range str { // 舍弃第二个参数
		fmt.Printf("str[%d]=%c\n",i,str[i])
	}

}
