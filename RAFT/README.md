# RAFT
This directory is a copy of the repository by chapin666  
(https://github.com/chapin666/simple-raft)  
With the comments in raft.go translated to english

# simple implementation of raft consensus algorithm

---

### TODO
- [x] Leader election
- [x] Log replication

### Branch
- master： Leader election & Log replication
- election：Leader election
- replication：Log replication

### build
```
go build -o simple-raft
```

### run
```
go get go get github.com/mattn/goreman
goreman start
```

