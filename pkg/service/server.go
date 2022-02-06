package service

import "grpc-streams/pkg/pb"

type MyServer struct {
	pb.UnimplementedMyServiceServer
}

func New() *MyServer {
	myServer := &MyServer{}
	return myServer
}
