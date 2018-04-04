package main

import (
	"fmt"
	"net"
	"log"
	"os"
)

func CheckError(err error) {
	if err!=nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func main() {
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}

//处理
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			Log(conn.RemoteAddr().String(), "connection error: ",err)
			return
		}

		Log(conn.RemoteAddr().String(), "recive data string:\n", string(buffer[:n]))
	}
}