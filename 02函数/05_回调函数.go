package main

import "fmt"

func main() {
	a := callFunc(1,2,add2)
	fmt.Println(a)
}

func callFunc(a,b int,ftest FuncType) (ret int){
	fmt.Println("callFunc")
	ret = ftest(a,b)
	return
}

func add2(a,b int) int {
	return a+b
}

type FuncType func(int,int) int