package main

import (
	"fmt"
	"time"
)

type person3 struct {
	id   int
	name string
}

func main() {

	now := time.Now()      //当前时间
	year := now.Year()     //当前年
	month := now.Month()   //当前月
	day := now.Day()       //当前日
	hour := now.Hour()     //当前小时
	second := now.Second() //当前秒
	//时间戳
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	//将时间戳转换为具体时间  使用time.Unix
	currTime := time.Unix(timeStamp1, 0)
	fmt.Println(currTime)
	fmt.Println(timeStamp1, timeStamp2)
	fmt.Println(year, month, day, hour, second)

	//时间间隔
	//time.Sleep(5 * time.Second)
	//fmt.Println("睡眠结束")

	//时间计算
	//add
	after := now.Add(1 * time.Minute) //计算一个小时后的时间
	before := after.Sub(now)          //after-now
	fmt.Println(after, before)

	//定时器,每隔多少秒触发
	//tick := time.Tick(time.Second)
	//f := x()
	//for tmp := range tick {
	//	fmt.Println(tmp)
	//	fmt.Println(f())
	//	if tmp.After(after) {
	//		break
	//	}
	//}

	//时间格式化 2006-1-2-15-4-5
	format := now.Format("2006年1月2日 - 15时4分5秒.000毫秒")
	fmt.Println(format)

	//解析字符串类型的时间
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	timestr := "2003-01-15 21:00:00"
	//转换为自定义的时间类型
	//注意这里需要转换的字符串与parse中layout参数格式得一样
	timeobj1, err := time.Parse("2006-1-2 15:4:5", timestr)
	if err != nil {
		fmt.Println("日期转换失败")
		return
	}
	fmt.Println(timeobj1)
	//设置时区,调用ParseInLocation方法,有三个参数
	timeobj2, err := time.ParseInLocation("2006-1-2 15:4:5", timestr, location)
	if err != nil {
		fmt.Println("设置时区失败")
		return
	}
	fmt.Println(timeobj2)
	timeobj3 := timeobj2.Format("2006/1-2 15:4:5")
	fmt.Println(timeobj3)

	//p := person3{
	//	id:   1,
	//	name: "123",
	//}

}

func x() func() int {
	i := 0
	fun := func() int {
		i++
		return i
	}
	return fun
}
