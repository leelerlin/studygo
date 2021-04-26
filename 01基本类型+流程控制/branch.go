package main

import (
	"fmt"
	"io/ioutil"
)

func main() { // 左括号必需与函数名同行
  if1()
  if2()
}

func if1 () {
	const filename = "./abc.txt";
    contents ,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n",contents)
		fmt.Println(123)
	}
}

func if2 () {
	const filename = "abc.txt";
	if contents ,err := ioutil.ReadFile(filename);err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n",contents)
		fmt.Println(123)
	}
}
