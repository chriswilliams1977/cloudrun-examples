#gRPC Overview
#gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types.
#By default, gRPC uses protocol buffers as the Interface Definition Language (IDL) for describing both the service interface
#and the structure of the payload messages.

#gRPC lets you define four kinds of service method
#Unary RPCs where the client sends a single request to the server and gets a single response back
#Server streaming RPCs where the client sends a request to the server and gets a stream to read a sequence of messages back
#Client streaming RPCs where the client writes a sequence of messages and sends them to the server, single response back
#Bidirectional streaming RPCs where both sides send a sequence of messages using a read-write stream

#Protoc generates client and server APIs
#gRPC users typically call these APIs on the client side and implement the corresponding API on the server side.
#On the server side, the server implements the methods declared by the service and runs a gRPC server to handle client calls.
#The gRPC infrastructure decodes incoming requests, executes service methods, and encodes service responses.
#On the client side, the client has a local object known as stub (for some languages, the preferred term is client) that implements the same methods as the service.
#The client can then just call those methods on the local object, wrapping the parameters for the call in the appropriate protocol buffer message type
#- gRPC looks after sending the request(s) to the server and returning the server’s protocol buffer response(s).

#The client and the servers talk to each other, over HTTP/2 (transport layer on gRPC).
#The messages are binary data(thanks to Protobuf), but the communication is in plaintext. Fortunately, gRPC has SSL/TLS integration,
#that can be used to authenticate the server (from the client’s perspective), and to encrypt message exchanges.
#3 types of connection
#insecure connection, all data transfered between client and server is not encrypted.
#Server-side TLS -  data is encrypted, but only the server needs to provide its TLS certificate to the client. You can use this type of connection if the server doesn’t care which client is calling its API.
#Mutual TLS - We use it when the server also needs to verify who’s calling its services. So in this case, both client and server must provide their TLS certificates to the other.
#Cloud Run now forces HTTPS on any application running on *.run.app domain. Therefore you must use port 443 to connect to a gRPC server.

#In order to validate the identity of the Server (authenticate it), the client uses the certification authority (CA) certificate
#to authenticate the CA signature on the server certificate. You can provide the CA certificate to your client or
#rely on a set of trusted CA certificates included in your operating system (trusted key store).

#a context.Context object, containing your data, and that will exist during all the lifetime of the request;

#If GO modules not loading check GOLAND project setting to see if GO modules are enabled
install_protoc:
	go get -u github.com/golang/protobuf/protoc-gen-go
build_proto:
	protoc \
        --proto_path api/v1 \
        --go_out "plugins=grpc:pkg/api/v1" \
        message.proto
build_image:
	docker build -t grpc-ping .