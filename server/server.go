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
var messageStreams []chan proto.ServerMessage

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
	messageStreams = make([]chan proto.ServerMessage, 0)

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
func (c *Server) ConnectClient(in *proto.Connection, stream proto.ChittyChat_ConnectClientServer) error {
	msgChan := make(chan proto.ServerMessage)
	messageStreams = append(messageStreams, msgChan)
	for {
		msg := <-msgChan
		log.Println("Sending message!")
		stream.Send(&msg)
	}
	return nil
}

func (c *Server) DisconnectClient(ctx context.Context, in *proto.Disconnection) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (c *Server) SendClientMessage(ctx context.Context, in *proto.ClientMessage) (*proto.Empty, error) {
	log.Println(in.Message)
	for i := 0; i < len(messageStreams); i++ {
		messageStreams[i] <- proto.ServerMessage{
			Username:  "user",
			Message:   in.Message,
			Timestamp: 0,
		}
	}
	return &proto.Empty{}, nil
}
