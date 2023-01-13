package main

import (
	"fmt"
	"sync"
	"time"
)

var rwlock sync.RWMutex
var wg3 sync.WaitGroup
var z int = 0
var lock sync.Mutex //互斥锁

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg3.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go write()
	}
	wg3.Wait()
	fmt.Println("总时间为", time.Now().Sub(start))
	/*
		总结:
		1.在读占比大的场景下,使用读写锁的效率更高
		2.如果读和写相差不大,拿互斥锁和读写锁相差不大
	*/
}
func read() {
	rwlock.RLock() //读写锁中读操作:RLock()
	time.Sleep(1)
	rwlock.RUnlock()
	wg3.Done()
}
func write() {
	rwlock.Lock()
	z += 1
	time.Sleep(10)
	rwlock.Unlock()
	wg3.Done()
}
