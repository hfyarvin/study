package main

import (
	pb "../../proto" // 引入proto包
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	// "google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	grpcTest1()
}

func Getclient(conn *grpc.ClientConn) pb.HelloClient {
	c := pb.NewHelloClient(conn)
	return c
}

func grpcTest1() {
	conn, err := GetConnect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := Getclient(conn)

	//请求
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"

	//响应
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Println("err:", err)
	}

	if r != nil {
		log.Println(r.Message)
	}
}

func GetConnect() (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile("../../key/server.pem", "tgq")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
