package main

import (
	"errors"
	"fmt"
)

func foo(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("除数为0")
		result = 10
		return
	}
	result = a / b
	return
}

func main() {
	m, n := 2111, 0
	if p, err := foo(m, n); err == nil {
		fmt.Println("商为", p)
	} else {
		fmt.Println("发生了异常", err)
		fmt.Println(p)
	}
}
