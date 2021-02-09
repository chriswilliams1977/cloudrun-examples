package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/GoogleCloudPlatform/golang-samples/run/grpc-ping/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/status"
)

//THIS IS THE PING SERVICE CLIENT
////IMPLEMENTS PingServiceClient interface

type pingService struct {
	pb.UnimplementedPingServiceServer
}

func (s *pingService) Send(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("sending ping response")
	return &pb.Response{
		Pong: &pb.Pong{
			Index:      1,
			Message:    req.GetMessage(),
			ReceivedOn: ptypes.TimestampNow(),
		},
	}, nil
}

func (s *pingService) SendUpstream(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if conn == nil {
		return nil, fmt.Errorf("no upstream connection configured")
	}

	p := &pb.Request{
		Message: req.GetMessage() + " (relayed)",
	}

	hostWithoutPort := strings.Split(os.Getenv("GRPC_PING_HOST"), ":")[0]
	tokenAudience := "https://" + hostWithoutPort
	resp, err := PingRequest(conn, p, tokenAudience, os.Getenv("GRPC_PING_UNAUTHENTICATED") == "")
	if err != nil {
		log.Printf("PingRequest: %q", err)
		c := status.Code(err)
		return nil, status.Errorf(c, "Could not reach ping service: %s", status.Convert(err).Message())
	}

	log.Print("received upstream pong")
	return &pb.Response{
		Pong: resp.Pong,
	}, nil
}