package main

import "fmt"

func main() {
	s := "golang"
	fmt.Println(s)
	// 字符串本质：是不可修改的byte数组
	for i, ele := range s {
		fmt.Println(i, ele)
	}
}
