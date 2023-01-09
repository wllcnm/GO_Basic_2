package main

import "fmt"

type hobby []string

// Student 匿名字段就是结构体中嵌套结构体
type Student struct {
	Class //匿名字段,student默认包含了Class的所有字段
	id    int
	name  string
}
type Class struct {
	id    int
	cid   int
	cname string
}

// 匿名类型不只有struct作为字段,自定义类型、内置类型都可以作为匿名字段
type test struct {
	int //支持内置类型
	float64
	hobby //支持自定义类型
}

func main() {
	student := Student{id: 101, name: "孙笑川"}
	student.Class = Class{id: 102, cid: 20, cname: "快乐家园"}
	fmt.Println(student.Class.id)
	//修改班级信息
	student.Class.cname = "快乐星球"
	fmt.Println(student)

	t := test{}
	t.int = 7
	t.float64 = 3.14
	t.hobby = []string{"123", "456"}
	fmt.Println(t)
}
