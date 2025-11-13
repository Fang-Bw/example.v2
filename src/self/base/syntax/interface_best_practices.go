package main

import "fmt"

func main() {
	t := T{data: "Hello, World!"}
	fmt.Println(t.read())
	t.write("Hello, Go!")
	fmt.Println(t.read())

	m1 := map[int]T{1: {"1"}}
	fmt.Println(m1[1].read())
	// 这里会报错，因为m1[1]是T类型，不是*T类型，即使是T类型，也不能取地址，因为go限制map中的元素不能取地址，所以需要使用指针类型
	//m1[1].write("2")
	fmt.Println(m1[1].read())

	m2 := map[int]*T{1: {"1"}}
	fmt.Println(m2[1].read())
	m2[1].write("2")
	fmt.Println(m2[1].read())
}

type T struct {
	data string
}

func (t T) read() string {
	return t.data
}

func (t *T) write(data string) {
	t.data = data
}
