package main

import (
	pb "../../proto"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/grpclog"
)

const (
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
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
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
		log.Println(err)
	}
	if r != nil {
		log.Println(r.Message)
	}
}

/**
 * http://localhost:50051/debug/requests
 * http://localhost:50051/debug/events
 */
