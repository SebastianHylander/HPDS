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
var messageStreams map[int]chan proto.ServerMessage
var time int64

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
	messageStreams = make(map[int]chan proto.ServerMessage, 0)
	time = 0

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
	if in.Timestamp > time {
		time = in.Timestamp
	}
	time++

	msgChan := make(chan proto.ServerMessage)
	messageStreams[int(in.ClientId)] = msgChan
	users[int(in.ClientId)] = in.Username

	go sendMessageToChannels(proto.ServerMessage{
		Username:  users[int(in.ClientId)],
		Message:   "Joined the chat!",
		Timestamp: time,
	})

	_, chanExists := messageStreams[int(in.ClientId)]

	for chanExists {
		msg := <-msgChan
		stream.Send(&msg)
		_, chanExists = messageStreams[int(in.ClientId)]
	}
	return nil
}

func (c *Server) DisconnectClient(ctx context.Context, in *proto.Disconnection) (*proto.Empty, error) {
	if in.Timestamp > time {
		time = in.Timestamp
	}
	time++

	for _, channel := range messageStreams {
		channel <- proto.ServerMessage{
			Username:  users[int(in.ClientId)],
			Message:   "Left the chat",
			Timestamp: time,
		}
	}

	delete(messageStreams, int(in.ClientId))
	delete(users, int(in.ClientId))
	return &proto.Empty{}, nil
}

func (c *Server) SendClientMessage(ctx context.Context, in *proto.ClientMessage) (*proto.Empty, error) {
	if in.Timestamp > time {
		time = in.Timestamp
	}
	time++

	sendMessageToChannels(proto.ServerMessage{
		Username:  users[int(in.ClientId)],
		Message:   in.Message,
		Timestamp: time,
	})
	return &proto.Empty{}, nil
}

func sendMessageToChannels(message proto.ServerMessage) {
	for _, channel := range messageStreams {
		channel <- message
	}
}
