package main

import (
	"fmt"
	"strconv"
)

// 0 - h
// 1 - i
// 2 -
// 3 - 乱码
// 4 - 乱码
// 5 - 乱码
// 6 - p
func forStr() {
	str := "hi 我p"
	// 直接遍历乱码
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d - %c\n", i, str[i])
	}
}

// 0 - h
// 1 - i
// 2 -
// 3 - 我
// 6 - p
func forStr2() {
	str := "hi 我p"
	// 能正常的打印，但是后面的索引就是按照字节进行标注的
	for i, s := range str {
		fmt.Printf("%d - %c\n", i, s)
	}
}

// 0 - h
// 1 - i
// 2 -
// 3 - 我
// 4 - p
func forStr3() {
	str := "hi 我p"
	// 能正常的打印，一切符合预期
	for i, s := range []rune(str) {
		fmt.Printf("%d - %c\n", i, s)
	}
}

func main() {
	forStr()
	a := 60
	c := strconv.FormatInt(int64(a), 2)
	b := strconv.FormatInt(int64(a), 10)
	fmt.Println(a, c, b)
}
