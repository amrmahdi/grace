# grace

1. Start server in on terminal: `go run server/main.go`
2. Start client in second terminal: `go run client/main.go`
3. Send SIGINT to server <kbd>Ctrl+C</kbd>
4. Doesn't wait for RPC to finish. Client receives error.

### Output

#### Server

```
$ go run server/main.go
Starting server...
Starting sleep
^CShutting down gracefully...
Error:  accept tcp [::]:5543: use of closed network connection
```

#### Client

```
$ go run client/main.go
2017/03/08 14:10:24 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: dial tcp :5543: getsockopt: connection refused"; Reconnecting to {:5543 <nil>}
Error:  rpc error: code = 13 desc = transport is closing
```
