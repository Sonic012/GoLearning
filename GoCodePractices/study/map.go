package main

import "fmt"

func main() {
	var m map[string]int
	//m=make(map[string]int,100)
	m = map[string]int{"a": 1, "b": 2, "c": 3}
	m["d"] = 10 // 添加新的值
	m["a"] = 4  // 修改已有的值
	fmt.Println(m["a"])
	delete(m, "a")                   // 删除map中的某个key
	fmt.Println(m["a"])              // 返回int默认值0
	if v, exists := m["a"]; exists { // m["a"]会拆包成两个值，一个是元素值，另一个是判断这个key是否存在，分号后面是判断条件，前面是初始化
		fmt.Println(v)
	} else {
		fmt.Println("map中不存在a这个key")
		fmt.Println(exists) // false

	}

	for k := range m {
		fmt.Println(k)
	}
}
