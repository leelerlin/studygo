package main

import "fmt"

func main() {
	a:=10
	str := "mike"

	// 匿名函数，没有函数名
	// 方法1
	f1 := func(){
		fmt.Println("a=",a) // 获取外部变量
		fmt.Println("str=",str)
	}
	f1()
	// 方法2，起别名后调用
	type FuncType func()
	var f2 FuncType
	f2 = f1
	f2()
	// 方法3 ,定义匿名函数同时调用
	func(){
		fmt.Printf("a=%d,str=%s\n",a,str)
	}()

	// 带参数的匿名函数
	func(i,j int){
		fmt.Printf("i=%d,j=%d\n",i,j)
	}(10,20)

	// 有参数有返回值
	max,min := func(i,j int) (max,min int){
		if i > j {
			max = i
			min = j
		}else{
			max = j
			min = i
		}
		return
	}(30,40)
	fmt.Printf("max is %d,min is %d",max,min)

	// 递归求和
	sum := sum(100)
	fmt.Printf("sum 1~100 %d\n",sum)

	fiboRet := fibo(5)
	fmt.Printf("fiboRet=%d\n",fiboRet)
}

// 递归求和
func sum(i  int ) (int) {
	if(i == 1){
		return 1
	}
	return  i + sum(i-1)
}

// 斐波那契  求第n个数的值  1，1，2，3，5
func fibo(n int ) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
