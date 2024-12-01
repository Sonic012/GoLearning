package test

import (
	"fmt"
	"reflect"
	"rsc.io/quote"
)

func printType(v interface{}) {
	t := reflect.ValueOf(v)
	fmt.Println(t, t.Kind())
}

func test() {
	//var x int = 100
	//var y float64 = 3.14
	//var z string = "Go"

	// 获取变量的类型和种类
	//printType(x)
	//printType(y)
	//printType(z)
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}
