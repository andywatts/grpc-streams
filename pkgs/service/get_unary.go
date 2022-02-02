package service

import (
	"context"
	"fmt"
	"myprotos/pkgs/pb"
)

func (s *MyServer) GetUnary(ctx context.Context, myUnaryRequest *pb.MyUnaryRequest) (*pb.MyUnaryResponse, error) {
	response := fmt.Sprintf("%s redux", myUnaryRequest.Value)
	return &pb.MyUnaryResponse{Value: response}, nil
}
