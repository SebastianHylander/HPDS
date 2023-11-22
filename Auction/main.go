package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/SebastianHylander/HPDS/Auction/Node"
)

func main() {
	// Read the file
	file, err := os.Open("nodes.txt")
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

	// Print the nodes
	for i := range nodes {
		// make a Node from node.go with the given ip and port
		// call the function to start the node

		// Split the string into ip and port
		ipAndPort := strings.Split(nodes[i], " ")
		port, _ := strconv.Atoi(ipAndPort[1])

		newNode := Node.NewNode(ipAndPort[0], port)

		go newNode.Start(nodes)
	}
	for {
	}
}
