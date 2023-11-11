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
	ip         string
	port       int
	neighbour  proto.MutualExclusionClient
	hasToken   bool
	wantAccess bool
}

func NewNode(ip string, port int) *Node {
	node := &Node{
		ip:   ip,
		port: port,
	}
	return node
}

func (n *Node) Start(neighbourip string, neighbourport int, token bool) {
	n.hasToken = token
	go n.startServer()

	neighbour, err := n.connectToNeighbour(neighbourip, neighbourport)
	if err != nil {
		log.Fatalf("Could not connect to neighbour: %v", err)
	}

	n.neighbour = neighbour

	if n.hasToken {
		n.neighbour.HandoverToken(context.Background(), &proto.Token{})
	}

	go n.run()

}

func (n *Node) run() {
	for {
		// generate a random integer between 0 and 10000 (0 and 10 seconds)
		i := rand.Intn(10000)
		// sleep for the random amount of time
		time.Sleep(time.Duration(i) * time.Millisecond)
		n.wantAccess = true

		log.Print(n.port, "wants access to the file")

		for !n.hasToken {
		}

		log.Print(n.port, "got the token")

		// Write 'hello' at the buttom of the file output.txt
		file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Could not open file: %v", err)
		}
		_, err = file.WriteString(n.ip + ":" + strconv.Itoa(n.port) + " got the token and accessed the file\n")
		if err != nil {
			log.Fatalf("Could not write to file: %v", err)
		}
		defer file.Close()

		log.Print(n.port, "wrote to file and handed over the token")

		n.wantAccess = false
	}

}

func (n *Node) HandoverToken(ctx context.Context, in *proto.Token) (*proto.Empty, error) {
	n.hasToken = true
	for n.wantAccess {
	}
	n.hasToken = false
	n.neighbour.HandoverToken(ctx, in)

	return &proto.Empty{}, nil
}

func (n *Node) startServer() {
	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(n.port))
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", n.port)

	// Register the grpc server and serve its listener
	proto.RegisterMutualExclusionServer(grpcServer, n)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func (n *Node) connectToNeighbour(ip string, port int) (proto.MutualExclusionClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to port %d", port)
	} else {
		log.Printf("%d Connected to the server at port %d\n", n.port, port)
	}
	return proto.NewMutualExclusionClient(conn), nil
}
