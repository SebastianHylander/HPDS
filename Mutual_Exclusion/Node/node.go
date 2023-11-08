package main

import (
	"flag"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Node struct {
	id         int64
	portNumber int
}

var (
	nodePort   = flag.Int64("cPort", 0, "client port number")
	serverPort = flag.Int("sPort", 0, "server port number (should match the port used for the server)")
)

func main() {
	// Create a node
	node := &Node{
		id:         *nodePort,
		portNumber: *serverPort,
	}
	waitForMessage(node)
}

func waitForMessage(node *Node) {

}

func connectToServer() (proto.MutualExclusionNode, error) { //Hvordan laver vi denne
	// Dial the server at the specified port.
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *serverPort)
	} else {
		log.Printf("Connected to the server at port %d\n", *serverPort)
	}
	return proto.NewMutualExclusionClient(conn), nil //Hvordan laver vi denne
}
