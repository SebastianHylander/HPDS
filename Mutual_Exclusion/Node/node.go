package main

import (
	"log"
	"net"
	"strconv"

	proto "github.com/SebastianHylander/HPDS/Mutual_Exclusion/gRPC"
	"google.golang.org/grpc"
)

type Node struct {
	proto.UnimplementedMutualExclusionServer
	ip        string
	port      int
	neighbour *proto.MutualExclusionClient
}

func Start(ip string, port int, neighbourip string, neighbourport int) {

	neighbour, err := connectToNeighbour(neighbourip, neighbourport)
	if err != nil {
		log.Fatalf("Could not connect to neighbour: %v", err)
	}

	// Create a node
	node := &Node{
		ip:        ip,
		port:      port,
		neighbour: &neighbour,
	}
	waitForMessage(node)
}

func waitForMessage(node *Node) {

}

func startServer(node *Node) {
	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(node.port))
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", node.port)

	// Register the grpc server and serve its listener
	proto.RegisterMutualExclusionServer(grpcServer, node)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func connectToNeighbour(ip string, port int) (proto.MutualExclusionClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to port %d", port)
	} else {
		log.Printf("Connected to the server at port %d\n", port)
	}
	return proto.NewMutualExclusionClient(conn), nil
}
