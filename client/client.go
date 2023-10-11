package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	proto "simpleGuide/grpc"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id         int
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
	go waitForTimeRequest(client)

	for {

	}
}

func waitForTimeRequest(client *Client) {
	// Connect to the server
	serverConnection, _ := connectToServer()

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		log.Printf("Client asked for time with input: %s\n", input)

		clientStartTime := time.Now()

		// Ask the server for the time
		timeReturnMessage, err := serverConnection.AskForTime(context.Background(), &proto.AskForTimeMessage{
			ClientId: int64(client.id),
		})
		clientEndTime := time.Now()

		if err != nil {
			log.Printf(err.Error())
		} else {
			serverStartTime := time.Unix(timeReturnMessage.StartTime, 0)
			serverEndTime := time.Unix(timeReturnMessage.EndTime, 0)

			totalTimeElapsed := clientEndTime.Sub(clientStartTime)
			elapsedAtServer := serverEndTime.Sub(serverStartTime)
			transportTime := elapsedAtServer - totalTimeElapsed

			newClientTime := serverEndTime.Add(transportTime / 2)

			log.Printf("Time since message arrived to the server: %s\n", time.Since(serverStartTime))
			log.Printf("The time elapsed at the server was: %s\n", elapsedAtServer)
			log.Printf("The time used on transport was: %s\n", transportTime.String())
			log.Printf("The client's current time is: %s\n", time.Now().String())
			log.Printf("The client should change their time to: %s\n", newClientTime.String())

		}
	}
}

func connectToServer() (proto.TimeAskClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *serverPort)
	} else {
		log.Printf("Connected to the server at port %d\n", *serverPort)
	}
	return proto.NewTimeAskClient(conn), nil
}
