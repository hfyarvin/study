package main

import (
	"log"
	"net"
	"net/http"

	pb "../../proto" // 引入编译生成的包

	"golang.org/x/net/context"
	"golang.org/x/net/trace" // 引入trace包
	"google.golang.org/grpc"
	// "google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:50052"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "Hello" + in.Name + "."

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Println(err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServer(s, HelloService)

	go startTrace()

	log.Println("Listen on" + Address)

	s.Serve(listen)
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	go http.ListenAndServe(":50051", nil)
	log.Println("Trace Listen on 50051")
}
