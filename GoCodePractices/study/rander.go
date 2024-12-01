package main

import (
	"fmt"
	randv2 "math/rand/v2"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d    ", randv2.IntN(100)) // 这里的IntN里面维护一个全局rander，某次生成的种子是不固定的，所以每次生成的随机数也随机
	}
}

func TestRandSeed(t *testing.T) {
	source := randv2.NewPCG(123, 456)
	for i := 0; i < 5; i++ {
		source.Seed(123, 456)
		rander := randv2.New(source)
		fmt.Printf("%d  ", rander.IntN(100)) // 60  56  81  94  43
	}
}

func main() {
	TestRandSeed(nil)
}
