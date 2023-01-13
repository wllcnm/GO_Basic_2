package main

import (
	"fmt"
	"sync"
)

var m = make(map[int]int)
var wg4 = sync.WaitGroup{}
var m2 = sync.Map{}

// 并发安全的map
//
//	func main() {
//		for i := 0; i < 20; i++ {
//			wg4.Add(1)
//			go func(i int) { //map中设置键值对
//				set(i, i+100)
//				fmt.Println("key:", i, "value:", get(i))
//			}(i)
//		}
//		wg4.Wait()
//	}
func main() {
	for i := 0; i < 20; i++ {
		go func(i int) {
			wg4.Add(1)
			m2.Store(i, i+100) //通过安全的map存储k和v
			value, _ := m2.Load(i)
			fmt.Printf("key%v , value%v \n", i, value)
			wg4.Done()
		}(i)
	}
	wg4.Wait()
}
func get(key int) int {
	return m[key]
}
func set(key int, value int) {
	m[key] = value
}
