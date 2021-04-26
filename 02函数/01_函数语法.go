package main

import "fmt"

func main() {
	firstFunc()
	argFunc(123,1)

	var a int
	a = argRet(456 ,1)
	fmt.Println(a)

	argsFunc(1)
	argsFunc(1,2)
	argsFunc(1,2,3)

	// 不定参数传递
	test(1,2,3,4)

	// 无参有返回值
	a = argHasRet()
	fmt.Println(a)
}


// 无参数无返回
func firstFunc(){
	fmt.Println("do firstFunc")
}

// 有参数无返回
func argFunc(num int,num2 int){
	fmt.Printf("%d\n",num + num2)
}

// 有参数有返回
func argRet(num , num2 int)  int{
	return num + 1
}

// 不定参数  如多个参数此参数必需写在最后一个参数
func argsFunc(args ...int){
	/*for i:=0;i<len(args);i++{
		fmt.Printf("args[%d]=%d\n",i,args[i])
	}*/
	fmt.Println("不定参数>>>")
	// 利用迭代遍历数组
	for i,data := range(args){
		fmt.Printf("args[%d]=%d\n",i,data)
	}
	fmt.Println("===============")
}

// 不定参数传递
func test(args ...int){
	argsFunc(args[:2]...) // 0~1 传递
	argsFunc(args[2:]...) // 2~最后
}

// 无参有返回值
func argHasRet() (ret int) {
	ret = 2
	return
}