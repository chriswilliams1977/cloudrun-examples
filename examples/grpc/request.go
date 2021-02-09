package main

// [START cloudrun_grpc_request]
// [START run_grpc_request]

import (
	"context"
	"time"

	pb "github.com/GoogleCloudPlatform/golang-samples/run/grpc-ping/pkg/api/v1"
	"google.golang.org/grpc"
)

// pingRequest sends a new gRPC ping request to the server configured in the connection.
func pingRequest(conn *grpc.ClientConn, p *pb.Request) (*pb.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := pb.NewPingServiceClient(conn)
	return client.Send(ctx, p)
}

// [END run_grpc_request]
// [END cloudrun_grpc_request]

// PingRequest creates a new gRPC request to the upstream ping gRPC service.
func PingRequest(conn *grpc.ClientConn, p *pb.Request, url string, authenticated bool) (*pb.Response, error) {
	if authenticated {
		return pingRequestWithAuth(conn, p, url)
	}
	return pingRequest(conn, p)
}