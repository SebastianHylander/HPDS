# HPDS

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
The package that we have used for our solution is the net package. Packages such as net provides us a portable interface with tools, in this case network i/o (tcp/ip).
The data structure we use is byte arrays.

b) Does your implementation use threads or processes? Why is it not realistic to use threads?
Our implementation uses processes as we run on host, and therefore not threads. It is not realistic to use threads, because the protocol does not run across a network, which means they run locally (no networking errors).

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
Re-ordering would be handled as our solution is connection-oriented - using TCP, we sequentially increment a number for each packet we send in order to tell the host in which order the different packages are to be read. This will also tell whether something is lost or duplicated. To ensure acknowledgement, we increment seq between the client and server (3-way handshake).

d) In case messages can be delayed or lost, how does your implementation handle message loss?
Our solution is a 3-way handshake, as the homework description (2) is all about, and therefore it does not contain message loss, as it is not required. However, to implement message we would handle errors, for example by timeouts between the client and the host, since every package is numbered the client can wait for a response from the host for every package sent. If the client does not get a response in a defined timeout-time the package is sent again. If this causes a duplication on the host's side the host will only process one of the packages since they have the same number.

e) Why is the 3-way handshake important?
The 3-way handshake is important because it allows us to reliably transmit data between devices.
First, the client sends a synchronization (SYN) request to the receiving device to establish a connection.
Second, the server acknowledges the SYN request by sending a SYN-ACK message back, which also contains its own initial sequence number.
Third, the client acknowledges the SYN-ACK message, and the connection is established.
Generally, It is reliable because it requeres acknowledgement each step: the client and server agrees on the sequence numbers, which are sent and acknowledged.
