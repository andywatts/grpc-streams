package main

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
}

func main() {
	logger.Log.Info("Starting")

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8080))
	Check(err)
	myServer := server.New()

	grpcOpts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(middleware.CtxTagsUnary, middleware.LogUnary, middleware.AuthUnary),
		grpc_middleware.WithStreamServerChain(middleware.CtxTagsStream, middleware.LogStream, middleware.AuthStream),
	}
	grpcServer := grpc.NewServer(grpcOpts...)
	pb.RegisterMyServiceServer(grpcServer, myServer)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

	fmt.Println("End")
}
