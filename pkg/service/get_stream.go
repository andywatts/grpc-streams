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
type World string
type User string
type Chunk string

var (
	worldClients map[*World]map[*User]*Connection
	clientChunk  map[*User]*Chunk
)

func init() {
	worldClients = make(map[*World]map[*User]*Connection)
	SetInterval(PingClients, 4000)
}

func (s *MyServer) GetStream(stream pb.MyService_GetStreamServer) error {
	in, _ := stream.Recv()
	user := stream.Context().Value("user").(User)
	world := World(in.Value)

	s.Subscribe(&world, &user, &stream)

	for {
		in, err := stream.Recv()
		if err != nil {
			log.Printf("GetStream receive error %s", err.Error())
			Unsubscribe(&world, &user)
			return err
		}

		/*
			switch evt := in.Event.(type) {
			case *pb.MyStreamRequest_World:
			case *pb.MyStreamRequest_Chunk:
			}
		*/

		log.Printf("GetStream got message %s from %s", in.Value, user)
		resp := &pb.MyStreamResponse{
			Event: &pb.MyStreamResponse_Message{Message: &pb.Message{
				Value: fmt.Sprintf("Server received %s from %s", in.Value, user),
			}},
		}
		if err := stream.Send(resp); err != nil {
			log.Printf("GetStream send response error %s", err.Error())
			return err
		}
	}
	//return <-conn.error
}

func Unsubscribe(world *World, user *User) {
	// Delete user from world
	delete(worldClients[world], user)

	// Delete world if empty
	if len(worldClients[world]) == 0 {
		delete(worldClients, world)
	}
}

func PingClients() {
	logger.Sugar.Infof("Pinging %d worlds", len(worldClients))
	resp := &pb.MyStreamResponse{
		Event: &pb.MyStreamResponse_Message{Message: &pb.Message{
			Value: fmt.Sprintf("Server time %s", time.Now()),
		}},
	}

	for world, clients := range worldClients {
		for client, conn := range clients {
			logger.Sugar.Infof("GetStream#PingClients World: %s  Client: %s", world, client)
			conn.stream.Send(resp)
		}
	}
}

func (s *MyServer) Subscribe(world *World, user *User, stream *pb.MyService_GetStreamServer) {
	// Create world if needed
	if _, ok := worldClients[world]; !ok {
		worldClients[world] = make(map[*User]*Connection)
	}
	// Store stream in world
	conn := Connection{stream: *stream}
	worldClients[world][user] = &conn
}
