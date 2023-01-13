package main

import "fmt"

// 通道为引用类型,不需要传指针
// 可以在chan的左右加上<-表明,该变量是只写或只读的
func f1(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	//需要使用make初始化的类型:slice map chan
	var ch = make(chan int, 1) //通道初始化,如果不设置缓冲区会导致死锁
	//发送值给通道
	ch <- 10
	//从通道中取值
	x := <-ch
	fmt.Println(x)

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)
	for i := range ch2 {
		fmt.Println(i)
	}

}
