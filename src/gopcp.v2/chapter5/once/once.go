package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var count int
	var once sync.Once
	max := rand.Intn(100)
	// 虽然在一个循环里面，但是once只会执行一次，最后 零值++ 得到1
	for i := 0; i < max; i++ {
		once.Do(func() {
			count++
		})
	}
	fmt.Printf("Count: %d.\n", count)
}
