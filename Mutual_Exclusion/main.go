package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	ip         string
	portNumber int
}

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
	var nodes []Node

	// Read the file line by line
	for scanner.Scan() {
		// Get the line
		line := scanner.Text()

		// Split the line into a slice of strings
		split := strings.Split(line, " ")

		// Convert the strings to the correct types
		ip := split[0]

		port, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("Could not convert port to int: %v", err)
		}

		// Create a node
		node := Node{
			ip:         ip,
			portNumber: port,
		}

		// Add the node to the slice
		nodes = append(nodes, node)
	}

	// Print the nodes
	for _, node := range nodes {
		// make a Node from node.go with the given ip and port

	}
}
