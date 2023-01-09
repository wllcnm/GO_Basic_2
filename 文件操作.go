package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//通过xxx.read()读取
	//readall()

	//通过bufio.NewReader,每次读取一行
	//readByBufio()

	//通过os.ReadFile读取,可一次性读取完,但存在内存,不易读大文件内容
	//readByIOUtil()

	//读取本地电影文件
	//openmovie()

	//写操作
	//write()

	//写操作通过BufIo
	writeByBufIo()

	//写操作通过os.WriteFile
	writeByIOUtil()
}
func read() {
	//打开文件
	txtObj, err := os.Open("./txt/test.txt") //相对路径
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	//读取文件
	var b = make([]byte, 128)
	n, err := txtObj.Read(b) //read函数,返回值为读取到的字节数和err
	if err != nil {
		fmt.Println("读取失败")
	}
	fmt.Println(string(b))
	fmt.Println(n) //n为读取到的字节数
	//关闭文件
	defer txtObj.Close()
}
func readall() {
	//打开文件
	txtObj, err := os.Open("./txt/test.txt") //相对路径
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	var tmpfile []byte
	var nsum int
	for {
		var tmp = make([]byte, 128)
		n, err := txtObj.Read(tmp) //read函数,返回值为读取到的字节数和err
		if err == io.EOF {         //EOF End Of File 读到文件末尾
			//fmt.Println(string(b[:]))
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			break
		}
		tmpfile = append(tmpfile, tmp...)
		nsum += n
		//fmt.Println(string(tmp))
		//fmt.Println(string(b[:]))
		//fmt.Println(n) //n为读取到的字节数
	}
	fmt.Println(string(tmpfile))
	fmt.Println(len(tmpfile))
	fmt.Println(nsum)
	//关闭文件
	defer txtObj.Close()
}

func readByBufio() {
	file2, err := os.Open("./txt/test.txt")
	if err != nil {
		fmt.Println("打开失败")
	}
	defer func(file2 *os.File) {
		err := file2.Close()
		if err != nil {
			fmt.Println("文件关闭失败")
		}
	}(file2)
	newReader := bufio.NewReader(file2)
	for {
		readString, err := newReader.ReadString('\n') //以\n作为分隔符,每次只会读取一行
		if err != nil {
			return
		}
		fmt.Println(readString)
	}
}

// os.readFile()使用方便
func readByIOUtil() {
	content, err := os.ReadFile("./txt/test.txt")
	if err != nil {
		println("读取失败")
		return
	}
	fmt.Println(string(content))
}
func openmovie() {
	//打开一部电影
	movie, err := os.Open("G:\\电影\\阳光电影www.ygdy8.com.神秘海域.2022.HD.1080P.中英双字.mkv")
	if err != nil {
		return
	}
	defer movie.Close()
	var movie1 []byte
	for {
		var tempmov = make([]byte, 1024)
		_, err := movie.Read(tempmov)
		if err == io.EOF {
			//fmt.Println(tempmov)
			break
		}
		if err != nil {
			break
		}
		movie1 = append(movie1, tempmov...)
		//fmt.Println(tempmov)
	}
	fmt.Println(string(movie1))
}

// 写入文件
func write() {
	fileobj, err := os.OpenFile("./txt/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer fileobj.Close()
	str := "追加的字符串"
	fileobj.Write([]byte(str)) //将str字符串转换为[]byte类型
	fileobj.WriteString("hello,world")
}
func writeByBufIo() {
	fileobj, err := os.OpenFile("./txt/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer fileobj.Close()
	newWriter := bufio.NewWriter(fileobj)
	newWriter.WriteString("蔡徐坤很美\n") //将内容写入缓冲区
	newWriter.Flush()                //将内容写入磁盘
}

func writeByIOUtil() {
	str := "溧阳徐志摩"
	//WriteFile()方法是默认覆盖写入的
	err := os.WriteFile("./txt/test.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println("写入失败")
		return
	}

}
