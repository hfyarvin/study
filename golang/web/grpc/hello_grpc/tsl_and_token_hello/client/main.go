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
	Address = "127.0.0.1:50052"

	OpenTLS = true
)

//自定义认证
type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am ke",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	if OpenTLS {
		return true
	}

	return false
}

func GetConnect() (*grpc.ClientConn, error) {
	var err error
	var opts []grpc.DialOption //认证
	if OpenTLS {
		creds, err := credentials.NewClientTLSFromFile("../../key/server.pem", "tgq")
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	//自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))
	conn, err := grpc.Dial(Address, opts...) //...表示所有数组都作为参数
	if err != nil {
		return nil, err
	}
	return conn, err
}

func test() {
	conn, err := GetConnect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Println(err)
	}
	if r != nil {
		log.Println(r.Message)
	}
}

func main() {
	test()
}
