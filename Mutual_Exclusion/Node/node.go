package Node

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	proto "github.com/SebastianHylander/HPDS/Mutual_Exclusion/gRPC"
	"google.golang.org/grpc"
)

type Node struct {
	proto.UnimplementedMutualExclusionServer
	ip        string
	port      int
	neighbour proto.MutualExclusionClient
}

var wantAccess bool
var hasToken bool
var node *Node

func Start(ip string, port int, neighbourip string, neighbourport int, token bool) {

	wantAccess = false
	hasToken = token

	// Create a node
	node = &Node{
		ip:        ip,
		port:      port,
		neighbour: nil,
	}
	startServer()

	neighbour, err := connectToNeighbour(neighbourip, neighbourport)
	if err != nil {
		log.Fatalf("Could not connect to neighbour: %v", err)
	}

	node.neighbour = neighbour

	go run()

	for {
	}
}

func run() {
	// generate a random integer between 0 and 100000 (0 and 100 seconds)
	i := rand.Intn(100000)
	// sleep for the random amount of time
	time.Sleep(time.Duration(i) * time.Millisecond)
	wantAccess = true

	for !hasToken {
	}

	// Write 'hello' at the buttom of the file output.txt
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	_, err = file.WriteString("hello\n")
	if err != nil {
		log.Fatalf("Could not write to file: %v", err)
	}
	defer file.Close()

	wantAccess = false

	go run()

}

func HandoverToken(ctx context.Context, in *proto.Token) (*proto.Empty, error) {

	hasToken = true
	for wantAccess {
	}
	hasToken = false
	node.neighbour.HandoverToken(ctx, in)

	return &proto.Empty{}, nil
}

func startServer() {
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
