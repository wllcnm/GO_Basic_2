package main

import (
	"fmt"
	reflect2 "reflect"
)

type student struct {
	Name  string `json:"name" ini:"s_name"`   //`json:"name" ini:"s_name"`为Tag
	Score int    `json:"score" ini:"s_score"` //`json:"score" ini:"s_score"`为Tag
}

// /想要反射调用方法方法名必须大写
func (s student) GetUp() {
	fmt.Println("起床")
}

func (s student) Sleep() {
	fmt.Println("睡觉")
}

func (s student) Eat(peo string) {
	fmt.Println(peo, "正在吃饭")
}

// 结构体反射
func main() {
	stu1 := student{
		Name:  "cxk",
		Score: 100,
	}
	//通过反射获取结构体所有字段信息
	Type := reflect2.TypeOf(stu1)
	fmt.Println(Type.Name(), Type.Kind()) //student struct
	//遍历结构体变量的所有字段,NumField为结构体中变量的个数
	for i := 0; i < Type.NumField(); i++ {
		filedObj := Type.Field(i) //Filed()方法为获取变量的元信息,传入的是序列号
		fmt.Println("Name为", filedObj.Name, "Type为", filedObj.Type, "Tag为", filedObj.Tag)
		fmt.Println(filedObj.Tag.Get("json"), filedObj.Tag.Get("ini"))
	}
	//根据指定名字获取结构体字段
	filedObj2, ok := Type.FieldByName("Name")
	if ok {
		fmt.Println(filedObj2.Type, filedObj2.Name, filedObj2.Tag)
	} else {
		fmt.Println("没有找到指定的值")
	}
	//获取结构体中的方法数量
	methodNum := printMethodNum(stu1)

	//循环调用结构体中的方法
	for i := 0; i < methodNum; i++ {

	}

}

func printMethodNum(a any) int {
	typeOf := reflect2.TypeOf(a)

	value := reflect2.ValueOf(a)

	NumMethod := value.NumMethod() //获取变量的方法数量

	fmt.Println(NumMethod)

	for i := 0; i < NumMethod; i++ {
		methodType := typeOf.Method(i).Type
		//拿到方法的名称
		fmt.Println("method name", typeOf.Method(i).Name)
		//拿到方法的类型
		fmt.Println("method", methodType)

		//通过反射调用方法传递的参数必须为[]reflect.value
		var args []reflect2.Value

		/*
			无法获取方法或函数的参数名称.
			这是因为名称对于调用方法或函数的人来说并不重要.重要的是参数的类型及其顺序.
		*/
		//如何拿到方法中参数的类型
		typeName := value.Method(i).Type().Name()
		fmt.Println(typeName)
		//有参数方法的调用
		if typeOf.Method(i).Name == "Eat" {
			of := reflect2.ValueOf("jojo")
			var x = []reflect2.Value{of}
			value.Method(i).Call(x)
		} else {
			//value值调用方法,调用者必须为value类型,用type不行
			value.Method(i).Call(args)
		}

	}
	return NumMethod
}
