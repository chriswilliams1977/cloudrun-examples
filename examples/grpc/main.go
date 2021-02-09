// Sample grpc-ping acts as an intermediary to the ping service.
package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/GoogleCloudPlatform/golang-samples/run/grpc-ping/pkg/api/v1"
)

// [START cloudrun_grpc_server]
// [START run_grpc_server]
func main() {
	log.Printf("grpc-ping: starting server...")

	//Get env vars
	//Can be passed in at command line PORT=9090 or in code
	port := os.Getenv("PORT")
	//if no port set to 8080
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	//Creates a TCP listener on port you want
	//gRPC uses HTTP/2, which multiplexes multiple calls on a single TCP connection. All gRPC calls over that connection go to one endpoint
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}


	//creates a new gRPC server with a server service which can be called via an API
	//attach the Ping service to the server
	//Remember server implements service interface to create API that can be called - PingServiceServer interface
	//RegisterService registers a service and its implementation to the gRPC server. - Server API ready for calls
	grpcServer := grpc.NewServer()
	pb.RegisterPingServiceServer(grpcServer, &pingService{})
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

// [END run_grpc_server]
// [END cloudrun_grpc_server]

// conn holds an open connection to the ping service.
var conn *grpc.ClientConn

//init func is called before main
//SET UP NEW GRPC CONNECTION
func init() {
	if os.Getenv("GRPC_PING_HOST") != "" {
		var err error
		conn, err = NewConn(os.Getenv("GRPC_PING_HOST"), os.Getenv("GRPC_PING_INSECURE") != "")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Starting without support for SendUpstream: configure with 'GRPC_PING_HOST' environment variable. E.g., example.com:443")
	}
}