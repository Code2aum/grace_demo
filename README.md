# gRPC Currency Server

This repository contains a gRPC server implemented in Go for managing currency values. The server allows you to perform CRUD (Create, Read, Update, Delete) operations on currency values using the gRPC protocol.

## Getting Started
### Usage

#### Running the server

To run the gRPC server:

```bash
go run server.go
```

#### Running the Client

To run the gRPC Client:

Create a currency:
```bash
go run client.go create USD 1.0
```

Read a currency:
```bash
go run client.go read USD
```

Update a currency:
```bash
go run client.go update USD 1.2
```

Delete a currency:
```bash
go run client.go delete USD
```
