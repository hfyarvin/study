package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "git.vodjk.com/go-grpc/example/proto"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)
