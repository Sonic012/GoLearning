package main

import "fmt"

func fooFunc() int {
	a, b := 3, 4
	c := a + b
	defer fmt.Println("111", c) // defer 的注册
	fmt.Println(c)
	c = 222
	defer fmt.Println("222", c) // defer 后注册，先执行

	defer func() {
		defer fmt.Println("444")
		defer fmt.Println("555")
		fmt.Println(333, c)

	}()
	c = 100
	return c

}

type person struct {
	Name string
}

func (p *person) SetName() {
	fmt.Println(p.Name)
}

func main() {
	p := person{Name: "张三"}
	defer p.SetName()
	p = person{Name: "李四"}
}
