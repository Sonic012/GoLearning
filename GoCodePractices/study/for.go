package main

import "fmt"

func main() {
	i := 0
	//for ; i < 3; i++ { // 这里的第一个分号不能省
	//	fmt.Println(i)
	//	// i++等价于把之前的i++放在这里。分号不能省
	//}

	for { // 这里是无限循环
		i++
		if i%2 == 0 {
			continue // 满足条件时跳过循环体，继续下一次循环
		}
		fmt.Println(i)
		if i > 10 {
			break // 满足条件时退出循环的条件
		}
	}
}
