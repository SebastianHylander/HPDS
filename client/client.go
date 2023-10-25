package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	proto "github.com/SebastianHylander/HPDS/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id         int
	username   string
	portNumber int
}

var (
	clientPort = flag.Int("cPort", 0, "client port number")
	serverPort = flag.Int("sPort", 0, "server port number (should match the port used for the server)")
)

func main() {
	// Parse the flags to get the port for the client
	flag.Parse()

	// Create a client
	client := &Client{
		id:         1,
		portNumber: *clientPort,
	}

	// Wait for the client (user) to ask for the time
	go waitForMessage(client)

	for {

	}
}

func waitForMessage(client *Client) {
	// Connect to the server
	serverConnection, _ := connectToServer()

	stream, _ := serverConnection.ConnectClient(context.Background())

	go logMessages(stream)

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		// Ask the server for the time
		serverConnection.SendClientMessage(context.Background(), &proto.ClientMessage{
			ClientId:  int64(client.id),
			Message:   input,
			Timestamp: 0,
		})

	}
}

func connectToServer() (proto.ChittyChatClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *serverPort)
	} else {
		log.Printf("Connected to the server at port %d\n", *serverPort)
	}
	return proto.NewChittyChatClient(conn), nil
}

func logMessages(stream proto.ChittyChat_ConnectClientClient) {
	for {
		message, _ := stream.Recv()

		if message != nil {
			log.Print(message.Timestamp)
			log.Print(" - ")
			log.Print(message.Username)
			log.Print(" : ")
			log.Println(message.Message)
		}

	}
}
