package main

import "fmt"

func main() {
	s := make([]int, 3)
	replaceElement(s) // 切片的传递，复制slice header传递
	fmt.Println(s)    // 打印 [1 1 1]
}

func replaceElement(s []int) {
	s[0] = 1
	s1 := s[:] // 切片的赋值，复制slice header
	s1[1] = 1
	s2 := s[2:]
	s2[0] = 1 // 切片的子切片，复制slice header
}
