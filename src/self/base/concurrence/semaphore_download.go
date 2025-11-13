package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 每个时刻限制只能有3个goroutine的并发，用于获取10个资源
func main() {

	const (
		MaxResources   = 10
		MaxConcurrence = 3
	)

	semaphore := make(chan struct{}, MaxConcurrence)
	wg := sync.WaitGroup{}
	wg.Add(MaxResources)

	for i := 0; i < MaxResources; i++ {
		go semaphore_operate(semaphore, i, &wg)
	}
	wg.Wait()
}

func semaphore_operate(semaphore chan struct{}, goRoutineId int, wg *sync.WaitGroup) {
	// 通过向channel中写数据来表示获取资源（简洁且常用）；也可以通过在channel预置资源，对应此处改为获取数据
	semaphore <- struct{}{}
	fmt.Println(strconv.Itoa(goRoutineId) + "获取并打印了" + strconv.Itoa(goRoutineId))
	// 通过读channel中的数据来表示释放资源（简洁且常用）；也可以通过在channel预置资源，对应此处改为在channel中写数据
	<-semaphore
	wg.Done()
}
