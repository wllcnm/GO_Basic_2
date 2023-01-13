package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1000000) //计数器+1,1代表只支持运行一个goroutine方法
	//go printfoo() //开启了一个goroutine
	//go printgo()
	for i := 0; i < 1000000; i++ {
		go printgo(i)
	}
	//每次完成一个线程都会使计数器-1
	//fmt.Println("hello")
	wg.Wait() //等计数器为0时,结束main函数
}
func printfoo() {
	fmt.Println("foofoofoo")
	wg.Done() //计数器-1
}
func printgo(i int) {
	fmt.Println("gogogo", i)
	wg.Done() //计数器-1
}
