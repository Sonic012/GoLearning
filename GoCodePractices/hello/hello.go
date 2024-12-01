package hello

import "fmt"

type Builder struct {
	value string
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetValue(value string) *Builder {
	b.value = value
	return b
}

func (b *Builder) Build() string {
	return b.value
}

func hello() {
	result := Builder{}.SetValue("World").Build()
	fmt.Println(result) // 输出：{Hello}

}
