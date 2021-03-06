package service

import (
	"context"
	"fmt"
	"grpc-streams/pkg/pb"
)

func (s *MyServer) GetUnary(ctx context.Context, myUnaryRequest *pb.MyUnaryRequest) (*pb.MyUnaryResponse, error) {
	//response := fmt.Sprintf("Welcome user %d, got your message '%s'", ctx.Context["id"], myUnaryRequest.Value)
	//var jwt = ctx.Value("jwt").(jwt2.StandardClaims)

	response := fmt.Sprintf("Welcome user %s, got your message", ctx.Value("user").(string))
	return &pb.MyUnaryResponse{Value: response}, nil
}
