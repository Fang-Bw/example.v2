package main

import "fmt"

func main() {
	printNumbers()
}

// 交替打印1-10的奇数偶数
func printNumbers() {
	// 使用struct作为channel的元素类型，因为其大小是0字节，最节省内存
	odd := make(chan struct{})
	even := make(chan struct{})
	// 声明done用于主函数的routine等待两个go函数执行完成
	done := make(chan struct{})

	go func() {
		for i := 1; i <= 10; i += 2 {
			<-odd
			if i%2 != 0 {
				fmt.Println(i)
				even <- struct{}{}
			}
		}
	}()

	go func() {
		for i := 2; i <= 10; i += 2 {
			<-even
			if i%2 == 0 {
				fmt.Println(i)
				// 在打印10时，不能再向odd发送元素
				if i < 10 {
					odd <- struct{}{}
				} else {
					done <- struct{}{}
				}
			}
		}
	}()
	// 向odd中添加元素触发打印流程执行
	odd <- struct{}{}
	<-done
}
