package main

import (
	"fmt"
	"log"
	"net"

	pb "../../proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/metadata"    // 引入grpc meta包
	// "google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:50052"
)

type helloService struct{}

var HelloService = helloService{}

var authToken map[string]string

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	authToken, err := auth(ctx)
	if err != nil {
		return nil, err
	}

	resp := new(pb.HelloReply)
	resp.Message = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, authToken["appid"], authToken["appkey"])
	return resp, nil
}

func auth(ctx context.Context) (map[string]string, error) {
	authToken = make(map[string]string)
	//解析metada中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		fmt.Println("appid:", val)
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		fmt.Printf("appkey:", val)
		appkey = val[0]
	}
	if appid != "101010" || appkey != "i am key" {
		return nil, grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}

	authToken["appid"] = appid
	authToken["appkey"] = appkey
	return authToken, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	checkError(err)

	creds, err := credentials.NewServerTLSFromFile("../../key/server.pem", "../../key/server.key")
	checkError(err)

	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterHelloServer(s, HelloService)

	log.Println("Listen on " + Address + " with TLS + Token")

	s.Serve(listen)
}

func checkError(err error) {
	if err != nil {
		log.Println("err:", err.Error())
		panic(err)
	}
}
