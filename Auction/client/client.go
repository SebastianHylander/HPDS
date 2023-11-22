package Client

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	proto "github.com/SebastianHylander/HPDS/Auction/grpc"
	"google.golang.org/grpc"
)

type Client struct {
	port   int64
	bid    int
	server proto.AuctionSystemClient
}

var (
	clientPort = flag.Int64("port", 0, "The port of the client")
)

func main() {
	// Create a new client
	client := Client{port: *clientPort, bid: 0}

	// Connect to the server
	client.connectToServer()

	// Run the client
	client.run()
}

func (c *Client) connectToServer() {
	// Read the file
	file, err := os.Open("../nodes.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Create a slice to store the nodes
	var nodes []string

	// Read the file line by line
	for scanner.Scan() {
		// Get the line
		line := scanner.Text()

		// Add the node to the slice
		nodes = append(nodes, line)
	}

	lookingForServer := true

	for lookingForServer {
		line := nodes[rand.Intn(len(nodes))]

		// Split the line into a slice of strings
		split := strings.Split(line, " ")

		// Convert the strings to the correct types
		ip := split[0]

		port, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("Could not convert port to int: %v", err)
		}

		//connect to serverIp and port
		conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
		if err == nil {
			c.server = proto.NewAuctionSystemClient(conn)
			lookingForServer = false
		}

	}
}

func (c *Client) Makebid(bid int) {
	if bid > c.bid {
		c.bid = bid

		tryingToBid := true

		for tryingToBid {
			// Send the bid to the server
			_, err := c.server.MakeBid(context.Background(), &proto.Bid{Id: c.port, Amount: int64(bid)})

			if err != nil {
				c.connectToServer()
			} else {
				tryingToBid = false
			}
		}
	}
}

func (c *Client) GetResult() {

	gettingResult := true

	for gettingResult {
		// Get the result from the server
		result, err := c.server.GetResult(context.Background(), &proto.Empty{})

		if err != nil {
			c.connectToServer()
		} else {
			gettingResult = false
			fmt.Println(result.Result)
		}
	}
}

func (c *Client) run() {
	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		if input == "result" {
			go c.GetResult()
		} else {
			// Convert the input to an int
			bid, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Could not convert input to int")
			} else {
				go c.Makebid(bid)
			}
		}
	}
}
