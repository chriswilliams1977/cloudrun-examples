package main

// [START cloudrun_grpc_request_auth]
// [START run_grpc_request_auth]

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"

	pb "github.com/GoogleCloudPlatform/golang-samples/run/grpc-ping/pkg/api/v1"
)

// pingRequestWithAuth mints a new Identity Token for each request.
// This token has a 1 hour expiry and should be reused.
// audience must be the auto-assigned URL of a Cloud Run service or HTTP Cloud Function without port number.
func pingRequestWithAuth(conn *grpc.ClientConn, p *pb.Request, audience string) (*pb.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create an identity token.
	// With a global TokenSource tokens would be reused and auto-refreshed at need.
	// A given TokenSource is specific to the audience.
	tokenSource, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		return nil, fmt.Errorf("idtoken.NewTokenSource: %v", err)
	}
	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("TokenSource.Token: %v", err)
	}

	// Add token to gRPC Request.
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token.AccessToken)

	// Send the request.
	client := pb.NewPingServiceClient(conn)
	return client.Send(ctx, p)
}

// [END run_grpc_request_auth]
// [END cloudrun_grpc_request_auth]