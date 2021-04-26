package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
注释块
 */

// 初始化变量
func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("variableZeroValue() => %d %q \n",a,s)
}

// 赋值
func variableInitValue(){
	var a,b int = 3,4
	var s string  = "test"
	fmt.Printf("variableInitValue() => %d %d %s \n",a,b,s)
 }

// 不指定类型赋值变量,自动推导类型，必需有个初始值
func variabeType(){
	var a,b,c,d = 3,4,true,"def"
	b = 5
	fmt.Println("variabeType() =>",a,b,c,d)

}

// 函数内不使用var定义变量
func variabeShorter(){
	a,b,c := 4,5,"test"
	b = 6
	fmt.Println("variabeShorter() =>",a,b,c)
}

// 批量赋值
var (
	aa = 11
	bb = 22
	cc = "testt"
)

// 复数
func euler(){
	c := 3 + 4i
	fmt.Println( "euler() =>",cmplx.Abs(c))
}

// 强制转换
func triangle(){
	var a ,b int = 3,4
	var c int
	d := math.Sqrt(float64(a*a + b*b))
	c = int(d)
	fmt.Println("triangle() => ",c,d)
}

// 常量
func consts(){
	const filename string = "test.txt"
	const a,b = 3,4
	const (
		m = 1
		n = 2
	)
	var c int
	c = int(math.Sqrt(a*a+b*b)) // 常量会自动类型转换
	fmt.Println("consts() => ",c)
}

// 枚举(特殊的常量)
func enums(){
	// iota 常量自动生成器，每个一行，自动累加1
	// iota 是给常量赋值的
	const (
		php = iota //iota 枚举自增
		_  //缺省
		java
		python
		golang
	)
	fmt.Println("enums() =>",php,java,python,golang)

	const (
		kb = 1 << (iota*10)
		mb
		gb
	)
	fmt.Println("enums() =>",kb,mb,gb)
}

func main() {
	fmt.Println("Hello world~");
	variableZeroValue()
	variableInitValue()
	variabeType()
	variabeShorter()
	fmt.Println(aa,bb,cc)
	euler()
	triangle()
	consts()
	enums()
}
