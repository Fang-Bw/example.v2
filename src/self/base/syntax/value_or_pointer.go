package main

import "fmt"

type Receiver struct {
	Name   string
	Gender int
}

func (receiver Receiver) SetReceiverNameByValue() {
	receiver.Name = "1"
	fmt.Println("Inside SetReceiverNameByValue:", receiver.Name)
}

func (receiver *Receiver) SetReceiverNameByPointer() {
	receiver.Name = "1"
	fmt.Println("Inside SetReceiverNameByValue:", receiver.Name)
}

func main() {
	// 演示值接收器 - 不会修改原始值
	receiver := Receiver{Name: "original"}
	fmt.Println("Before SetReceiverNameByValue:", receiver.Name)
	receiver.SetReceiverNameByValue()
	fmt.Println("After SetReceiverNameByValue:", receiver.Name)

	fmt.Println("---")

	// 演示指针接收器 - 会修改原始值
	receiver2 := Receiver{Name: "original"}
	fmt.Println("Before SetReceiverNameByPointer:", receiver2.Name)
	receiver2.SetReceiverNameByPointer()
	fmt.Println("After SetReceiverNameByPointer:", receiver2.Name)
}
