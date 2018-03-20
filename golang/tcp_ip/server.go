package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":4040")
	checkErr(err)

	listen, err := net.ListenTCP("tcp", addr)
	checkErr(err)
	fmt.Println("Start server...")

	for {
		conn, err := listen.Accept()
		checkErr(err)
		go Handle(conn)
	}
}

const BufLength = 128

var users map[string]net.Conn = make(map[string]net.Conn, 10)

func Handle(conn net.Conn) {
	conn.Write([]byte("Welcome..."))

	for {
		data := make([]byte, 0)
		buf := make([]byte, BufLength)

		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				checkErr(err)
			}

			data = append(data, buf[:n]...)
			if n != BufLength {
				break
			}
		}

		cmd := strings.Split(string(data), "|")
		fmt.Println("命令:", cmd)

		switch cmd[0] {
		case "nick":
			fmt.Println("注册名称:" + cmd[1])
			users[cmd[1]] = conn
		case "say":
			for k, v := range users {
				if k != cmd[1] {
					fmt.Println("给" + k + "发送消息:" + cmd[2])
					v.Write([]byte(cmd[1] + ":[" + cmd[2] + "]"))
				}
			}
		}
	}
}
