package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

func main() {
	input := true
	for input {
		var line string
		fmt.Println("Do you want to host? (Write 'host')")
		fmt.Println("Do you want to connect? (Write 'connect')")
		fmt.Println("Do you want to exit? (Write 'exit')")
		fmt.Scanln(&line)

		if line == "host" {
			host()
			input = false
		} else if line == "connect" {
			getIp()
			input = false
		} else if line == "exit" {
			input = false
		} else {
			fmt.Println("Invalid input!")
		}
	}
}

func getIp() {
	var ip string
	fmt.Println("Write desired ip address (if empty the ip is 'localhost'):")
	fmt.Scanln(&ip)
	if ip == "" {
		ip = "localhost"
	}
	connect(ip)
}

func connect(ip string) {
	conn, err := net.Dial("tcp", ip+":8080")
	if err != nil {
		// handle error
	}
	seq := rand.Intn(10000)
	fmt.Println("Establishing connection sending seq '" + strconv.Itoa(seq) + "' to host")
	conn.Write([]byte(strconv.Itoa(seq)))
	output := strings.Split(read(conn), " ")

	seqreturn, _ := strconv.Atoi(output[0][2:])
	ack, _ := strconv.Atoi(output[1][2:])

	fmt.Println("Recieved seq '" + strconv.Itoa(seqreturn) + "' and ack '" + strconv.Itoa(ack) + "'")
	fmt.Println("Connection established! Sending message!")
	acknowlegde := "x=" + strconv.Itoa(seqreturn) + " y=" + strconv.Itoa(ack+1)
	write(conn, acknowlegde)
	write(conn, "Hello Host! I can feel a connection between us!")

}

func host() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			//handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	seq := read(conn)
	seqNumber, _ := strconv.Atoi(seq)
	fmt.Println("Received seq '" + seq + "' from client!")
	ack := strconv.Itoa(rand.Intn(10000))
	seqNumber = seqNumber + 1
	acknowlegde := "x=" + strconv.Itoa(seqNumber) + " y=" + ack
	fmt.Println("Sending ack '" + strconv.Itoa((seqNumber)) + "' and new seq '" + ack + "' back to the client!")
	write(conn, acknowlegde)
	output := strings.Split(read(conn), " ")

	ackreturn, _ := strconv.Atoi(output[0][2:])
	seqreturn, _ := strconv.Atoi(output[1][2:])

	fmt.Println("Recieved seq '" + strconv.Itoa(ackreturn) + "' and ack '" + strconv.Itoa(seqreturn) + "'")
	fmt.Print("Message from Client: ")
	fmt.Println(read(conn))
}

func read(conn net.Conn) string {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		// handle error
	}
	output := string(buffer[:n])

	return output
}

func write(conn net.Conn, message string) {
	conn.Write([]byte(message))
}
