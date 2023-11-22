package Node

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	proto "github.com/SebastianHylander/HPDS/Auction/grpc"
	"google.golang.org/grpc"
)

type Node struct {
	proto.UnimplementedAuctionSystemServer
	ip           string
	port         int
	id           int64
	leaderId     int64
	nodeId2Index map[int]int
	neighbour    int
	replications []proto.AuctionSystemClient
	bids         map[int64]int64
	inProgress   bool
	mapLock      sync.Mutex
}

func NewNode(ip string, port int) *Node {
	node := &Node{
		ip:           ip,
		id:           int64(port),
		port:         port,
		nodeId2Index: make(map[int]int),
		bids:         make(map[int64]int64),
	}
	return node
}

func (n *Node) Start(nodes []string) {
	go n.startServer()

	// Time to start the servers
	time.Sleep(10 * time.Second)

	for _, node := range nodes {
		// Split the string into ip and port
		ipAndPort := strings.Split(node, " ")
		port, _ := strconv.Atoi(ipAndPort[1])
		if !(port == n.port) {
			// make grpc connection to each node
			connection, err := n.connectToNeighbour(ipAndPort[0], port)
			if err != nil {
				log.Fatalf("Could not connect to node: %v", err)
			}
			n.replications = append(n.replications, connection)
		} else {
			n.neighbour = len(n.replications) - 1
			if n.neighbour < 0 {
				n.neighbour = len(nodes) - 1
			}
			n.replications = append(n.replications, nil)
		}
		// append id to the list of ids
		n.nodeId2Index[port] = len(n.nodeId2Index)

	}
	n.inProgress = true
	if n.replications[0] == nil {
		n.RunElection(context.Background(), &proto.ElectionStatus{ServerId: n.id})
	}
	go n.run()

}

func (n *Node) run() {
	log.Print("Starting auction")
	time.Sleep(60 * time.Second)
	n.inProgress = false
	log.Print("Auction finished and inProgress set to ", n.inProgress)
	for {
	}

}

func (n *Node) RunElection(ctx context.Context, in *proto.ElectionStatus) (*proto.Empty, error) {
	if in.ServerId == n.id {
		log.Print("Election done")
		n.leaderId = in.ServerId
		log.Print("Leader is: ", n.leaderId)
		return &proto.Empty{}, nil
	} else if in.ServerId < n.id {
		in.ServerId = n.id
	}
	n.leaderId = in.ServerId

	runElection := true
	for runElection {
		_, err := n.replications[n.neighbour].RunElection(ctx, in)
		if err != nil {
			n.neighbour = (n.neighbour - 1)
			if n.neighbour < 0 {
				n.neighbour = len(n.replications) - 1
			}
			if n.replications[n.neighbour] == nil {
				n.leaderId = n.id
				runElection = false
			}
		} else {
			runElection = false
		}

	}

	return &proto.Empty{}, nil
}

func (n *Node) MakeBid(ctx context.Context, in *proto.Bid) (*proto.Ack, error) {
	if n.inProgress {
		if n.leaderId == n.id {
			n.updateBid(in.Id, in.Amount)
			in.FromLeader = true
			for _, node := range n.replications {
				if node != nil {
					node.MakeBid(ctx, in)
				}
			}
		} else {
			if in.FromLeader {
				n.updateBid(in.Id, in.Amount)
			} else {
				_, err := n.replications[n.nodeId2Index[int(n.leaderId)]].MakeBid(ctx, in)
				if err != nil {
					n.RunElection(ctx, &proto.ElectionStatus{ServerId: n.id})
					n.MakeBid(ctx, in)
				}
			}
		}
	}
	return &proto.Ack{}, nil
}

func (n *Node) updateBid(bidderId int64, bidValue int64) {
	n.mapLock.Lock()
	n.bids[bidderId] = bidValue
	n.mapLock.Unlock()

}

func (n *Node) GetResult(ctx context.Context, in *proto.Empty) (*proto.Result, error) {
	if n.leaderId == n.id {
		if len(n.bids) == 0 {
			return &proto.Result{Result: "No bids"}, nil
		}
		highestBidder := int64(0)
		for bidder, bid := range n.bids {
			if bid > n.bids[highestBidder] {
				highestBidder = bidder
			}
		}
		var result string
		if n.inProgress {

			result = "Auction still in progress: Current winner is bidder " + strconv.Itoa(int(highestBidder)) + " with a bid of " + strconv.Itoa(int(n.bids[highestBidder]))
		} else {
			result = "Auction finished: Winner is bidder " + strconv.Itoa(int(highestBidder)) + " with a bid of " + strconv.Itoa(int(n.bids[highestBidder]))
		}
		return &proto.Result{Result: result}, nil
	} else {
		return n.replications[n.nodeId2Index[int(n.leaderId)]].GetResult(ctx, in)
	}
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
	proto.RegisterAuctionSystemServer(grpcServer, n)

	//30% chance of causing a simulated crash
	if rand.Intn(100) < 30 {
		go n.Crash(grpcServer)
	}

	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Println("Error serving: ", serveError)
	}
}

func (n *Node) connectToNeighbour(ip string, port int) (proto.AuctionSystemClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())

	return proto.NewAuctionSystemClient(conn), err
}

// Function to simulate a crash failure
func (n *Node) Crash(server *grpc.Server) {
	time.Sleep(time.Duration(20+rand.Intn(20)) * time.Second)
	log.Print("Crashing node with port: ", n.port)
	server.Stop()
}
