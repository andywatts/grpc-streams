package service

import (
	"fmt"
	"grpc-streams/pkg/logger"
	"grpc-streams/pkg/pb"
	. "grpc-streams/pkg/utils"
	"log"
	"time"
)

type Connection struct {
	stream pb.MyService_GetStreamServer
	error  chan error
}

var (
	roomClients map[string]map[string]Connection // roomClients[room][userId]connection
)

func init() {
	roomClients = make(map[string]map[string]Connection)
	SetInterval(PingClients, 4000)
}

func (s *MyServer) GetStream(stream pb.MyService_GetStreamServer) error {
	in, _ := stream.Recv()
	userId := stream.Context().Value("userId").(string)
	room := in.Value

	s.Subscribe(room, userId, &stream)

	for {
		in, err := stream.Recv()
		if err != nil {
			log.Printf("GetStream receive error %s", err.Error())
			Unsubscribe(room, userId)
			return err
		}

		log.Printf("GetStream got message %s from %s", in.Value, userId)
		resp := &pb.MyStreamResponse{
			Event: &pb.MyStreamResponse_Message{Message: &pb.Message{
				Value: fmt.Sprintf("Server received %s from %s", in.Value, userId),
			}},
		}
		if err := stream.Send(resp); err != nil {
			log.Printf("GetStream send response error %s", err.Error())
			return err
		}
	}
	//return <-conn.error
}

func Unsubscribe(room string, userId string) {
	// Delete user from room
	delete(roomClients[room], userId)

	// Delete room if empty
	if len(roomClients[room]) == 0 {
		delete(roomClients, room)
	}
}

func PingClients() {
	logger.Sugar.Infof("Pinging %d rooms", len(roomClients))
	resp := &pb.MyStreamResponse{
		Event: &pb.MyStreamResponse_Message{Message: &pb.Message{
			Value: fmt.Sprintf("Server time %s", time.Now()),
		}},
	}

	for room, clients := range roomClients {
		for client, conn := range clients {
			logger.Sugar.Infof("GetStream#PingClients Room: %s  Client: %s", room, client)
			conn.stream.Send(resp)
		}
	}
}

func (s *MyServer) Subscribe(room string, userId string, stream *pb.MyService_GetStreamServer) {
	// Create room if needed
	if _, ok := roomClients[room]; !ok {
		roomClients[room] = make(map[string]Connection)
	}
	// Store stream in room
	conn := Connection{stream: *stream}
	roomClients[room][userId] = conn

}
