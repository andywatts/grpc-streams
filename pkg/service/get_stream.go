package service

import (
	"fmt"
	"grpc-streams/pkg/pb"
	"io"
	"log"
	"time"
)

type Connection struct {
	stream pb.MyService_GetStreamServer
	error  chan error
}

var (
	clients map[string]Connection
)

func init() {
	clients = make(map[string]Connection)
}

func (s *MyServer) GetStream(stream pb.MyService_GetStreamServer) error {
	user_id := stream.Context().Value("user_id").(string)
	conn := Connection{stream: stream}
	clients[user_id] = conn

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Printf("GetStream eof")
			return nil
		}
		if err != nil {
			log.Printf("GetStream error %s", err.Error())
			return err
		}

		log.Printf("GetStream got message %s from %s", in.Value, user_id)
		resp := &pb.MyStreamResponse{
			Event: &pb.MyStreamResponse_ClientMessage{ClientMessage: &pb.MyStreamResponse_Message{
				Value: fmt.Sprintf("Server received %s from %s", in.Value, user_id),
			}},
		}
		if err := stream.Send(resp); err != nil {
			log.Printf("GetStream send response error %s", err.Error())
			return err
		}
	}
}

func PingClients(stream pb.MyService_GetStreamServer) {
	resp := &pb.MyStreamResponse{
		Event: &pb.MyStreamResponse_ClientMessage{ClientMessage: &pb.MyStreamResponse_Message{
			Value: fmt.Sprintf("Server time %s", time.Now()),
		}},
	}
	stream.Send(resp)
}
