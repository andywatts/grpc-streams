package service

import (
	"fmt"
	"grpc-streams/pkg/pb"
	"io"
	"log"
)

func (s *MyServer) GetStream(stream pb.MyService_GetStreamServer) error {
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

		log.Printf("GetStream got message %s", in.Value)
		resp := &pb.MyStreamResponse{
			Event: &pb.MyStreamResponse_ClientMessage{ClientMessage: &pb.MyStreamResponse_Message{
				Value: fmt.Sprintf("Server received %s", in.Value),
			}},
		}
		if err := stream.Send(resp); err != nil {
			log.Printf("GetStream send response error %s", err.Error())
			return err
		}
	}
}
