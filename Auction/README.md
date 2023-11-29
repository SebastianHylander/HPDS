To run our program, start be defining the nodes in nodes.txt in the following format
```
  [ip address] [port]
  [ip address] [port]
  [ip address] [port]
  [ip address] [port]
  [ip address] [port]
```
Where each line is a new node. The default file has 5 nodes all on local host

When nodes.txt is set run the auction using:
```
go run main.go
```
In the current directory (HPDS/Auction/)

This will start all the nodes each with their own gRPC-server

Then start the clients in their own terminals using:
```
go run client.go -port [int]
```
In the client directory (HPDS/Auction/client)

This will start a client and connect it to a random node from 'nodes.txt'

To make a bid as a client type any integer
To get the result of the auction 'result'
