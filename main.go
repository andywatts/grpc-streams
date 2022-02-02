package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"myprotos/pkgs/pb"
	server "myprotos/pkgs/service"
	"net"
)

var (
	client bool
)

func init() {
	flag.BoolVar(&client, "c", false, "run as client")
	flag.Parse()
}

func main() {
	if client {
		fmt.Println("Starting Client")
		return
	}

	fmt.Println("Starting Server")
	lis, _ := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	myServer := server.New()
	pb.RegisterMyServiceServer(grpcServer, myServer)
	grpcServer.Serve(lis)
}
