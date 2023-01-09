package main

import "fmt"

type Student1 struct {
	id     int
	name   string
	Class1 //匿名字段中的method,父类也可以调用
}
type Class1 struct {
	cid   int
	cname string
}
type fun1 func()

func main() {

	student1 := Student1{id: 10, name: "123"}
	student1.sayhi() //如果父类和子类都有同名方法,则优先调用父类的方法

	var f1 fun1
	f1.m1()
}

func (c *Class1) sayhi() {
	fmt.Println("hello,class")
}
func (s *Student1) sayhi() {
	fmt.Println("hello,student")
}
func (f fun1) m1() {
	fmt.Println("hello,funciton")
}
