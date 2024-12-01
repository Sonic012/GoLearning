package main

import "fmt"

type Human1 struct {
	Age    int
	Height float32
	Gender bool
}

func main() {
	var a Human1
	a = Human1{25, 166.9, false} // 初始化结构体实例。省略字段名时，值的顺序与结构体字段顺序一致，且类型需要对应一致
	fmt.Printf("%d  %.2f  %t\n", a.Age, a.Height, a.Gender)
	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a) // 展示了对应值和及其所属字段名
	fmt.Printf("%#v\n", a) // 展示了结构体实例的源代码(属于哪个包下的那个结构体)
}
