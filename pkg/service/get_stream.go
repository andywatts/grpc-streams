package service

import (
	"fmt"
	"grpc-streams/pkg/logger"
	"grpc-streams/pkg/pb"
	. "grpc-streams/pkg/utils"
	"io"
	"log"
	"time"
)

type Connection struct {
	stream pb.MyService_GetStreamServer
	error  chan error
}

var (
	roomClients map[string]map[string]Connection // roomClients[room][user_id]connection
)

func init() {
	roomClients = make(map[string]map[string]Connection)
	SetInterval(PingClients, 4000)
}

func (s *MyServer) GetStream(stream pb.MyService_GetStreamServer) error {
	in, _ := stream.Recv()
	user_id := stream.Context().Value("user_id").(string)
	room := in.Value

	// Create room if needed
	if _, ok := roomClients[room]; !ok {
		roomClients[room] = make(map[string]Connection)
	}

	// Store stream in room
	conn := Connection{stream: stream}
	roomClients[in.Value][user_id] = conn

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Printf("GetStream eof")
			RemoveClient(room, user_id)
			return nil
		}
		if err != nil {
			log.Printf("GetStream error %s", err.Error())
			RemoveClient(room, user_id)
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
	return <-conn.error
}

func RemoveClient(room string, user_id string) {
	// Delete user from room
	delete(roomClients[room], user_id)

	// Delete room if empty
	if len(roomClients[room]) == 0 {
		delete(roomClients, room)
	}
}

func PingClients() {
	logger.Sugar.Infof("Pinging %d rooms", len(roomClients))
	resp := &pb.MyStreamResponse{
		Event: &pb.MyStreamResponse_ClientMessage{ClientMessage: &pb.MyStreamResponse_Message{
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
