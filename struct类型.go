package main

import "fmt"

// 全局定义实体类:非main函数也可以使用
// 实体类型:定义类型用type,定义变量用var,定义常量用const
type person struct { //person的类型为struct
	id      int
	name    string
	address string
	age     int
	phone   string //电话号码用string来存储 int的范围为-2147483648~2147483647
}

func main() {

	//第一种赋值方式:不用按照顺序赋值
	p1 := person{id: 1, name: "jojo", address: "china", age: 18, phone: "18777355612"}
	fmt.Println(p1.age)

	//第二种赋值方式:按顺序赋值,不用写关键字
	p2 := person{1, "cxk", "china", 18, "18777355162"}
	fmt.Println(p2)
	//第三种赋值方式
	var p3 person
	p3.id = 1
	p3.name = "jojo"
	p3.address = "china"
	p3.age = 18
	p3.phone = "18777355162"
	fmt.Println(p2)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//比较两个人的年龄,返回年龄大的
	p4 := person{age: 20}
	p5 := person{age: 20}
	max := cmpAge(p4, p5)
	fmt.Println(max)

}

// 比较两个人的年龄
func cmpAge(a, b person) person {

	if a.age == b.age {
		panic("两个人年龄相同")
	} else if a.age > b.age {
		return a
	} else {
		return b
	}
}
