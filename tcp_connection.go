package main

import (
	"fmt"
	"net"
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
	//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	//status, err := bufio.NewReader(conn).ReadString('\n')
	//fmt.Println(status)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	fmt.Println(string(buffer[:n]))
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
	conn.Write([]byte("Hello Client"))
}
