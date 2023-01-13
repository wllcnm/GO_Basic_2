package main

import (
	"fmt"
	reflect2 "reflect"
)

type Person struct {
	id   int
	name string
}

func main() {
	/*
		1.如何判断变量类型?
		(1)通过断言
		(2)通过reflect.typeof,反射获得
	*/

	//基本数据类型
	f := 3.14
	reflectType(f) //类型为 float64 种类为 float64

	//结构体类型
	p := Person{id: 1, name: "jojo"}
	reflectType(p) //类型为 main.person 种类为 struct

	//切片类型
	var a []int
	var b []string
	reflectType(a) //类型为 []int 种类为 slice

	reflectType(b) //类型为 []string 种类为 slice

	//值类型信息
	var aa []int64
	aa = []int64{123, 345}
	reflectValue(aa)
	reflectSetValue(&aa)

	//reflectSetValue
	var bb int64
	bb = 1234
	//如果是int64类型的值就修改为100
	reflectSetValue(&bb)
	var cc string
	cc = "1234"
	reflectSetValue(&cc)
	//如果是string类型的值就修改为"这是string类型的值"
	fmt.Println(bb)
	fmt.Println(cc)

	//通过reflect.valueof得到的值类型为reflect.Value,要得到具体的类型需要调用reflect.valueof().kind
	valueOf := reflect2.ValueOf(1)
	valueOfKind := valueOf.Kind()
	fmt.Println(valueOfKind)
}

func reflectType(v any) {
	//反射就是在程序运行时,修改程序的状态
	type1 := reflect2.TypeOf(v)
	fmt.Println("类型为", type1)
	fmt.Println("种类为", type1.Kind())

}
func reflectValue(v any) {
	//value就是获取值,value.kind就是获取类型
	value := reflect2.ValueOf(v)
	kind := value.Kind()
	fmt.Println(value)
	fmt.Println(kind)
}
func reflectSetValue(v any) {
	//获取值
	value := reflect2.ValueOf(v)
	//获取类型,加上elem是为了传入指针类型的值时可以获取到原始的值,因为函数都说值拷贝,想要修改值必须传入指针
	kind := value.Elem().Kind()
	//通过类型来判断是否要改变值
	switch kind {
	case reflect2.Int64:
		value.Elem().SetInt(100)
	case reflect2.String:
		value.Elem().SetString("这是string类型")
	case reflect2.Slice:
		fmt.Println("这是切片类型")
	}
}
