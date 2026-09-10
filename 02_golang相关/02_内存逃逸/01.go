package main

import "fmt"

func foo() *int {
	x := 10
	return &x
}

func main() {
	ans := foo()
	fmt.Println(ans)
}
