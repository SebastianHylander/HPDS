# HPDS

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
The package that we have used for our solution is the net package. Packages such as net provides us a portable interface with tools, in this case network i/o (tcp/ip).
The data structure we use is byte arrays.

b) Does your implementation use threads or processes? Why is it not realistic to use threads?
Our implementation uses processes as we run on host, and therefore not threads. It is not realistic to use threads, because the protocol does not run across a network, which means they run locally (no errors).

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
Re-ordering would be handled as our solution is connection-oriented - using TCP, we use numbers for each packet we send in order to find out something is lost, duplicated or for reordering things. To ensure acknowledgement, we increment seq between the client and server (3-way handshake).

d) In case messages can be delayed or lost, how does your implementation handle message loss?
We handle message loss by our implementation being connection-oriented. The point of TCP is that we use numbers for each packet you send in order to find out something is lost, duplicated or for reordering things. So, how we use the seq of the various packets is the: 3-way handshake

e) Why is the 3-way handshake important?
The 3-way handshake is important because it allows us to reliably transmit data between devices.
First, the client sends a synchronization (SYN) request to the receiving device to establish a connection.
Second, the server acknowledges the SYN request by sending a SYN-ACK message back, which also contains its own initial sequence number.
Third, the client acknowledges the SYN-ACK message, and the connection is established.
Generally, It is reliable because it requeres acknowledgement each step: the client and server agrees on the sequence numbers, which are sent and acknowledged.
