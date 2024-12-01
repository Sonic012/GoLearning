package main

import "fmt"

/*
接口是一组行为规范的集合，它定义了某个对象可以做什么事情。
An interface is a collection of behavioral specifications
that define what an object can do.
*/

type Human interface {
	Say(int, int) int
}

func foo1(h Human) {
	c := h.Say(2, 4)
	fmt.Println(c)
}

type Tom struct {
}

func (Tom) Say(a, b int) int {
	return a - b
}

type Sean struct {
}

func (Sean) Say(a, b int) int {
	return a + b
}

func main() {
	var a Human
	var b Human
	//var d interface{} // 表示空的接口
	t := Tom{}
	a = t
	foo1(a)

	s := Sean{}
	b = s
	foo1(b)

}
