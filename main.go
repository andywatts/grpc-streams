package main

import (
	"context"
	"flag"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	fmt.Println("main init")
	flag.BoolVar(&clientMode, "c", false, "run as client")
	flag.Parse()
	//if err := logger.Init(logLevel); err != nil {
	//	fmt.Errorf("failed to initialize logger: %v", err)
	//}
}

func main() {
	ctx := context.Background()
	if clientMode {
		client.New().Run(ctx)
		return
	}
	logger.Log.Info("Starting")

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8080))
	Check(err)

	myServer := server.New()

	grpcOpts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(middleware.CtxTagsUnary, middleware.LogUnary),
		grpc_middleware.WithStreamServerChain(middleware.CtxTagsStream, middleware.LogStream),
	}
	grpcServer := grpc.NewServer(grpcOpts...)
	pb.RegisterMyServiceServer(grpcServer, myServer)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

	fmt.Println("End")
}
