package main

// [START cloudrun_grpc_conn]
// [START run_grpc_conn]

import (
	"crypto/tls"
	"crypto/x509"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewConn creates a new gRPC connection.
// host should be of the form domain:port, e.g., example.com:443
func NewConn(host string, insecure bool) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}

	//you don’t need to manually provide the Server certificate to your gRPC client in order to encrypt the connection.
	//
	//TLS is one of the authentication mechanisms that are built-in to gRPC. It has TLS integration and promotes the use
	//of TLS to authenticate the server, and to encrypt all the data exchanged between the client and the server
	//The X.509 v3 certificate format  - It encodes, among other things, the server’s public key and a digital signature (to validate the certificate’s authenticity).
	//
	//If you do NOT want to encrypt the connection, the Go grpc package offers the DialOption WithInsecure() for the Client.
	//This, plus a Server without any ServerOption will result in an unencrypted connection.
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		//SystemCertPool returns a copy of the system cert pool.
		//
		// On Unix systems other than macOS the environment variables SSL_CERT_FILE and
		// SSL_CERT_DIR can be used to override the system default locations for the SSL
		// certificate file and SSL certificate files directory, respectively. The
		// latter can be a colon-separated list.
		//
		//manually load the CA certs from the system with SystemCertPool().
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}

		//secure the connection
		//Package credentials implements various credentials supported by gRPC library, which encapsulate all the state needed by a client
		//to authenticate with a server and make various assertions,
		//e.g., about the client's identity, role, or whether it is authorized to make a particular call.
		//
		//With CA certificates included in the system (OS/Browser)
		//Use system certs  and tls config (tls.Config{}) will take care of loading your system CA certs.
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	return grpc.Dial(host, opts...)
}

// [END run_grpc_conn]
// [END cloudrun_grpc_conn]