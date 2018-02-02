package main

import (
	"net"

	pb "../../proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	// "google.golang.org/grpc/grpclog"
	"log"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService ...
var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

func main() {
	log.Println("This is the server main func")
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Println("failed to listen: %v", err)
	}

	//TLS
	// creds,err := credentials.NewServerTLSFromFile(certFile, serverNameOverride)
	creds, err := credentials.NewServerTLSFromFile("../../key/server.pem", "../../key/server.key")
	if err != nil {
		log.Println("Failed to generate credentials %v", err)
	}

	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))

	//注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	log.Println("Listen on " + Address + " with TLS")
	s.Serve(listen)
}
