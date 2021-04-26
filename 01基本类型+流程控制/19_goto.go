package main

import "fmt"

func main() {
	fmt.Println("abc")

	goto End

	fmt.Println("def")

	End:
		fmt.Println("fff")
}
