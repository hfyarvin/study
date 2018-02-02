package main

import (
	"net"

	pb "../../proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	// "google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

type helloService struct{}

var HelloService = helloService{}

//给接口定义与Service同名方法
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("server SayHello....")
	resp := new(pb.HelloReply)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServer(s, HelloService)

	log.Println("Listen on " + Address)

	s.Serve(listen)
}
