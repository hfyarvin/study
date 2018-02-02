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
