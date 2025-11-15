package main

import "fmt"

func appendTwice(s []int) {
	// 第一次 append
	s = append(s, 1)
	fmt.Println("第一次append的slice:", s) // 第一次append的slice: [1]
	// 第二次 append（基于第一次的结果）
	s = append(s, 2)
	fmt.Println("第二次append的slice:", s) // 第二次append的slice: [1 2]
}

func main() {
	s := make([]int, 0, 5) // len=0, cap=5
	appendTwice(s)
	fmt.Println("原始slice:", s) // 原始slice: []
}
