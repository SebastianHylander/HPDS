package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	proto "github.com/SebastianHylander/HPDS/Mutual_Exclusion/gRPC"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedMutualExclusionServer
	name string
	port int
}

// Used to get the user-defined port for the server from the command line
var port = flag.Int("port", 0, "server port number")

func main() {
	// Get the port from the command line when the server is run
	flag.Parse()

	// Create a server struct
	server := &Server{
		name: "serverName",
		port: *port,
	}
	go startServer(server)
	for {
	}
}

func startServer(server *Server) {
	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	proto.RegisterMutualExclusionServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func (s *Server) ConnectNode(in *proto.NodeConnection, stream proto.MutualExclusion_ConnectNodeServer) error {

	return nil
}
