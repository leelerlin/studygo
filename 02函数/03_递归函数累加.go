package main

func main() {
	sum := addFunc(100)
	println(sum)
}

func addFunc(i int) int{
	if i == 1 {
		return 1
	}

	return i + addFunc(i-1)
}
