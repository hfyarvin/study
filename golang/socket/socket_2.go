package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

func printConn(c net.Conn) {
	fmt.Println("Ip: ", c.RemoteAddr().String())
	fmt.Println("Address: ", c.LocalAddr())
	fmt.Println("Ip type: ", reflect.TypeOf(c.RemoteAddr().String()))
	fmt.Println("Address type: ", reflect.TypeOf(c.LocalAddr()))
}

func main() {
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	printConn(conn)

	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Size: ", n)

	buf := make([]byte, 10)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			conn.Close()
		}

		fmt.Print(string(buf[:n]))
	}

	fmt.Println(string(buf[:n]))
}
