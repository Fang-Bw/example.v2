package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 题目2：生产者-消费者模型
// 3个生产者，每个生产5个产品
// 2个消费者，消费这些产品
func main() {
	produceConsume()
}

func produceConsume() {
	// 长度为10的channel
	products := make(chan string, 10)
	// 用于表示生产者已经生产者完毕
	produceDone := sync.WaitGroup{}
	produceDone.Add(3)
	// 用于表示消费者已经消费完毕
	consumeDone := sync.WaitGroup{}
	consumeDone.Add(2)

	// 3个生产者，每个生产5个产品
	for i := 0; i <= 2; i++ {
		go Produce(products, "生产者"+strconv.Itoa(i), &produceDone)
	}
	// 2个消费者竞争性消费
	for i := 0; i <= 1; i++ {
		go Consume(products, "消费者"+strconv.Itoa(i), &consumeDone)
	}
	// 等待生产者生产完所有的产品之后，将通道标记为关闭，通道标记关闭，但是数据还是可供消费
	produceDone.Wait()
	close(products)
	// 等待消费者消费完所有的产品
	consumeDone.Wait()
}

func Produce(products chan string, producer string, produceDone *sync.WaitGroup) {
	for i := 0; i <= 4; i++ {
		productName := producer + "-产品" + strconv.Itoa(i)
		products <- productName
	}
	produceDone.Done()
}

func Consume(products chan string, consumer string, consumeDone *sync.WaitGroup) {
	for {
		productName, flag := <-products
		if !flag {
			break
		}
		fmt.Println(consumer + "消费-" + productName)
	}
	consumeDone.Done()
}
