package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	proto "github.com/SebastianHylander/HPDS/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id         int64
	username   string
	portNumber int64
}

var (
	clientPort = flag.Int64("cPort", 0, "client port number")
	serverPort = flag.Int("sPort", 0, "server port number (should match the port used for the server)")
	username   = flag.String("username", "", "Display name in chat")
)

var time int64

func main() {

	time = 0
	// Parse the flags to get the port for the client
	flag.Parse()

	// Create a client
	client := &Client{
		id:         *clientPort,
		username:   *username,
		portNumber: *clientPort,
	}

	// Wait for the client (user) to ask for the time
	waitForMessage(client)

}

func waitForMessage(client *Client) {
	// Connect to the server
	serverConnection, _ := connectToServer()

	stream, _ := serverConnection.ConnectClient(context.Background(), &proto.Connection{
		ClientId:  *clientPort,
		Username:  *username,
		Timestamp: time,
	})

	fmt.Println("Welcome to the chat!")
	fmt.Println("Write messages to other clients online!")
	fmt.Println("Leave the chat by writing 'disconnect'")

	go logMessages(stream)

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		time++

		if input == "disconnect" {
			serverConnection.DisconnectClient(context.Background(), &proto.Disconnection{
				ClientId:  int64(client.id),
				Timestamp: time,
			})
			break
		} else {
			serverConnection.SendClientMessage(context.Background(), &proto.ClientMessage{
				ClientId:  int64(client.id),
				Message:   input,
				Timestamp: time,
			})
		}

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
			if message.Timestamp > time {
				time = message.Timestamp
			}
			time++

			msgstring := strconv.Itoa(int(time)) + " - " + message.Username + ": " + message.Message
			log.Println(msgstring)
		}

	}
}
