package main

import (
	"fmt"
	"sync"
)

var (
	x     int        = 0
	wg1              = sync.WaitGroup{}
	lock1 sync.Mutex //互斥锁
)

func main() {
	wg1.Add(2)
	go add()
	go add()
	wg1.Wait()
	fmt.Println(x)
}
func add() {
	for i := 0; i < 5000; i++ {
		lock1.Lock() //加锁,加锁是为了保证在同一时间只有一个goroutine对变量操作
		x += 1
		lock1.Unlock() //释放锁
	}
	//fmt.Println(x)
	wg1.Done()
}
