package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i ++
		time.Sleep(time.Second)

		if i == 2 {
			continue // 跳过然后继续循环
		}

		if i == 5 {
			break // 跳出
		}
		fmt.Println(i,"\n")
	}
}
