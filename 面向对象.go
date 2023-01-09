package main

import (
	"fmt"
	"math"
)

type ages int
type Rectangle struct {
	height float64
	width  float64
}
type Circle struct {
	radius float64
}

func main() {

	r1 := Rectangle{
		height: 10,
		width:  20,
	}
	area1 := r1.area()
	fmt.Println(area1)
}

//method的语法如下：
//func (r ReceiverType) funcName(parameters) (results)
/*
在使用method的时候重要注意几点:
1.虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
2.method里面可以访问接收者的字段
3.调用method通过.访问，就像struct里面访问字段一样
*/
// 计算长方形面积
func (r Rectangle) area() float64 {
	r.width = 100
	return r.width * r.height
}

// 计算圆形面积
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
