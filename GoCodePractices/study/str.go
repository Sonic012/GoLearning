package main

import "fmt"

func main() {
	s := "Agolang"
	a := "A世"
	//c := s + " " + a + " world"
	for _, r := range s {
		//fmt.Printf("%d: %c (%[2]U)\n", i, r)
		fmt.Println(r)
	}
	fmt.Println(len(s))
	fmt.Printf("字节形式：% x\n", []byte(a))

}
