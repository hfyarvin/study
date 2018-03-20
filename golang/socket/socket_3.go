package main

import (
	"bufio"
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

	//将这个链接（connection）包装以下。将conn的内容都放入r中，但是没有进行读取，让步我们一会对其进行操作。
	r := bufio.NewReader(conn)
	for {
		//将r的内容也就是conn的数据按照换行符进行读取。
		line, err := r.ReadString('\n')
		if err == io.EOF {
			conn.Close()
		}
		fmt.Print(line)
	}
}

func printConn(c net.Conn) {
	fmt.Println("Ip: ", c.RemoteAddr().String())
	fmt.Println("Address: ", c.LocalAddr())
	fmt.Println("Data type: ", reflect.TypeOf(c.RemoteAddr().String()))
	fmt.Println("Address type: ", reflect.TypeOf(c.LocalAddr()))
}
