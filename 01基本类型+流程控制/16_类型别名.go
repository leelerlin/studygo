package main

import "fmt"

func main() {
	// 给 int64起个别名
	type bigint int64

	var a bigint
	fmt.Printf("a type is %T \n",a)

	type (
		long int64
		char byte
	)
	var b long = 11
	var c char = 'c'
	fmt.Printf("b is type %T,c is type %T \n",b,c)

}
