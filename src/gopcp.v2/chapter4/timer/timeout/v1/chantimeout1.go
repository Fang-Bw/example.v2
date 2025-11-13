package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 3)
		intChan <- 1
	}()
	select {
	// 执行到这里时，intChan没有数据可读，阻塞
	case e := <-intChan:
		fmt.Printf("Received: %v\n", e)
	// 执行到这里时，也会阻塞
	case <-time.NewTimer(time.Second * 5).C:
		fmt.Println("Timeout!")
	}
	//最后执行哪条语句，取决于谁阻塞的时间更短
}
