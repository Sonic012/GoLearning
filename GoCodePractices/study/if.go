package main

import "fmt"

func main() {
	//a,b:=10,true  // 一行可以同时声明多个变量并赋值
	if a, b := 10, true; a > 5 {
		fmt.Println("a >5")
	} else if b {
		fmt.Println("a <=5 and b is true")
	} else {
		fmt.Println("a <=5 and b is false")
	}
	fmt.Println(a) // 这里没法使用以上的a，变量，因为他们是块级作用域
}
