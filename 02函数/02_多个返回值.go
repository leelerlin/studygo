package main

import "fmt"

func main() {
	a,b,c := myFunc()
	fmt.Printf("a=%d,b=%d,c=%d\n",a,b,c)

	d,e,f := myFunc2()
	fmt.Printf("d =%d,e=%d,f=%d",d,e,f)
}

func myFunc() (int,int,int){
	return 1,2,3
}

// 推荐写法
func myFunc2() (a,b,c int){
	a,b,c = 1,2,3
	return
}