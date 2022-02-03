package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-streams/pkg/client"
	"grpc-streams/pkg/grpc/middleware"
	"grpc-streams/pkg/logger"
	"grpc-streams/pkg/pb"
	server "grpc-streams/pkg/service"
	. "grpc-streams/pkg/utils"
	"net"
)

var (
	clientMode bool
	logLevel   int = 0
)

func init() {
	flag.BoolVar(&clientMode, "c", false, "run as client")
	flag.Parse()
	if err := logger.Init(logLevel); err != nil {
		fmt.Errorf("failed to initialize logger: %v", err)
	}
}

func main() {
	ctx := context.Background()
	if clientMode {
		client.New().Run(ctx)
		return
	}

	fmt.Println("Starting Server")
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8080))
	Check(err)
	var opts []grpc.ServerOption
	opts = middleware.AddLogging(logger.Log, opts)
	grpcServer := grpc.NewServer(opts...)

	myServer := server.New()
	pb.RegisterMyServiceServer(grpcServer, myServer)
	grpcServer.Serve(lis)

	fmt.Println("End")
}
