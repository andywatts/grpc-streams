package client

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-streams/pkg/pb"
	. "grpc-streams/pkg/utils"
	"io"
	"log"
	"os"
	"time"
)

type client struct {
	pb.MyServiceClient
}

func New() *client {
	return &client{}
}

func (c *client) Run(ctx context.Context) {
	fmt.Println("Starting Client")

	// Connect
	connCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.DialContext(connCtx, "localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	Check(err)
	defer conn.Close()

	c.MyServiceClient = pb.NewMyServiceClient(conn)
	c.stream(ctx)
}

func (c *client) stream(ctx context.Context) {
	stream, err := c.MyServiceClient.GetStream(ctx)
	Check(err)
	defer stream.CloseSend()

	go c.receive(stream)
	c.send(stream)
}

func (c *client) receive(stream pb.MyService_GetStreamClient) error {
	for {
		res, err := stream.Recv()
		Check(err)

		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			log.Printf("stream canceled (usually indicates shutdown)")
			return nil
		} else if err == io.EOF {
			log.Printf("stream closed by server")
			return nil
		} else if err != nil {
			return err
		}

		switch evt := res.Event.(type) {
		case *pb.MyStreamResponse_Message:
			log.Printf("Response: %s", evt.Message.Value)
		case *pb.MyStreamResponse_ServerShutdown:
			log.Printf("The server is shutting down")
			return nil
		default:
			log.Printf("unexpected event from the server: %T", evt)
			return nil
		}
	}
}

func (c *client) send(stream pb.MyService_GetStreamClient) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	for {
		select {
		case <-stream.Context().Done():
			log.Printf("client send loop disconnected")
		default:
			if sc.Scan() {
				if err := stream.Send(&pb.MyStreamRequest{Value: sc.Text()}); err != nil {
					log.Printf("failed to send message: %v", err)
					return
				}
			} else {
				log.Printf("input scanner failure: %v", sc.Err())
				return
			}
		}
	}
}
