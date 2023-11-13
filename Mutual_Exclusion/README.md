## Mutual Exclusion

To run our program, start be defining the nodes in _nodes.txt_ in the following format

> [ip address] [port] \n
> [ip address] [port] \n
> [ip address] [port] \n
> [ip address] [port] \n
> [ip address] [port] \n

Where each line is a new node.
The default file has 5 nodes all on local host

When _nodes.txt_ is set run the program using

```
go run main.go
```

In the current directory (HPDS/MutualExclusion/)

The output of the program will be appended to _output.txt_, so make sure to clear the current lines in the file, if you want to start from new.
