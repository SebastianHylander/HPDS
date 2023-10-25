package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	proto "github.com/SebastianHylander/HPDS/grpc"

	"google.golang.org/grpc"
)

// Struct that will be used to represent the Server.
type Server struct {
	proto.UnimplementedChittyChatServer // Necessary
	name                                string
	port                                int
}

// Used to get the user-defined port for the server from the command line
var port = flag.Int("port", 0, "server port number")

var users map[int]string
var messageStreams []*grpc.ClientStream

func main() {
	// Get the port from the command line when the server is run
	flag.Parse()

	// Create a server struct
	server := &Server{
		name: "serverName",
		port: *port,
	}

	// Start the server
	go startServer(server)

	// Keep the server running until it is manually quit
	for {

	}
}

func startServer(server *Server) {

	users = make(map[int]string)
	messageStreams = make([]*grpc.ClientStream, 0)

	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	proto.RegisterChittyChatServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func ConnectClient(ctx context.Context, in *proto.Connection, opts ...grpc.CallOption) (proto.ChittyChat_ConnectClientClient, error) {
	users[int(in.GetClientId())] = in.GetUsername()

	/*log.Printf("Client with ID %d asked for the time\n", in.ClientId)
	startTime := time.Now().Unix()
	time.Sleep(time.Millisecond * 3498)
	endTime := time.Now().Unix()
	return &proto.TimeMessage{
		StartTime:  startTime,
		ServerName: c.name,
		EndTime:    endTime,
	}, nil
	*/

}

func (c *Server) DisconnectClient(ctx context.Context, in *proto.Disconnection) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (c *Server) SendClientMessage(ctx context.Context, in *proto.ClientMessage) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}
