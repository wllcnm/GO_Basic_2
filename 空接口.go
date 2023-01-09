package main

import "fmt"

// 空接口可以接受任何值
// 空接口的应用
// 1.空接口类型可以作为函数的参数
// 2.空接口可以作为map的value,类似java的object
func main() {
	var x interface{}
	x = 100
	x = "123"
	x = false
	x = func() {
		println("这是一个方法")
	}
	fmt.Println(x)
	reciver("123")
	//2.空接口可以作为map的value,类似java的object
	var m = make(map[string]interface{})
	m["id"] = 1
	m["name"] = "蔡徐坤"
	m["age"] = 18
	fmt.Println(m)

}

// 1.空接口类型可以作为函数的参数
func reciver(args ...interface{}) {
	fmt.Println(args)
}
