package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//等待客户连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed %v\n", err)
			continue
		}
		//启动一个单独的goroutine单独处理连接
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close() //延迟关闭连接
	//针对当前连接去处理接受信息
	for {
		newReader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := newReader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed,err is %v\n", err)
			return
		}
		recv := string(buf[:n])
		fmt.Printf("接受到的数据:%s", recv)
		conn.Write(buf[:n]) //将收到的数据返回给客户端
	}
}
