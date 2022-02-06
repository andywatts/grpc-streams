package service

import (
	"fmt"
	"grpc-streams/pkg/logger"
	"grpc-streams/pkg/pb"
	. "grpc-streams/pkg/utils"
	"log"
	"time"
)

var ()

func init() {
	SetInterval(PingClients, 4000)
}

func (s *MyServer) GetStream(stream pb.MyService_GetStreamServer) error {
	user := FindUser(&stream)
	for {
		in, err := stream.Recv()
		if err != nil {
			log.Printf("GetStream receive error %s", err.Error())
			user.Logout()
			return err
		}
		logger.Sugar.Infof("GetStream got %s from %s", in.Event, user.id)

		switch evt := in.Event.(type) {
		case *pb.MyStreamRequest_JoinWorld:
			domainName := DomainName(evt.JoinWorld.Value)
			world := FindWorld(domainName)
			world.Subscribe(user)

		case *pb.MyStreamRequest_ChangeChunk:
			user.chunk = Chunk(evt.ChangeChunk.Value)
		}

		resp := &pb.MyStreamResponse{
			Event: &pb.MyStreamResponse_Message{Message: &pb.Message{
				Value: fmt.Sprintf("Server received %s from %s", in.Event, user.id),
			}},
		}
		if err := stream.Send(resp); err != nil {
			log.Printf("GetStream send response error %s", err.Error())
			return err
		}
	}
	//return <-conn.error
}

func PingClients() {
	logger.Sugar.Infof("Pinging %d worlds and %d users", len(worlds), len(users))
	resp := &pb.MyStreamResponse{
		Event: &pb.MyStreamResponse_Message{Message: &pb.Message{
			Value: fmt.Sprintf("Server time %s", time.Now()),
		}},
	}

	for domainName, world := range worlds {
		logger.Sugar.Infof("GetStream#PingClients World: %s with: %d users", domainName, len(world.users))
		for _, user := range world.users {
			logger.Sugar.Infof("GetStream#PingClients World: %s user: %s", domainName, user.id)
			user.connection.stream.Send(resp)
		}
	}
}
