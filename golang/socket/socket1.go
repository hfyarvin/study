package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

func main() {
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	printConn(conn)

	//向服务端发送数据。用n接受返回的数据大小，用err接受错误信息。
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("向服务发送数据大小: ", n)

	buf := make([]byte, 1024)

	n, err = conn.Read(buf) //接收到内容的大小

	if err != nil && err != io.EOF { //io.EOF在网络编程中表示对端把链接关闭了。
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n])) //将接受的内容都读取出来
	conn.Close()
}

func printConn(c net.Conn) {
	fmt.Println("Ip: ", c.RemoteAddr().String())
	fmt.Println("Address: ", c.LocalAddr())
	fmt.Println("Data type: ", reflect.TypeOf(c.RemoteAddr().String()))
}
