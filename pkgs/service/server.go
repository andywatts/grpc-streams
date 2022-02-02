package service

import "myprotos/pkgs/pb"

type MyServer struct {
	ClientStreams map[string]chan *pb.MyStreamResponse
	pb.UnimplementedMyServiceServer
}

func New() *MyServer {
	myServer := &MyServer{
		ClientStreams: make(map[string]chan *pb.MyStreamResponse),
	}
	return myServer
}
