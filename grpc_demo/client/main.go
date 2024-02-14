package main

import (
	"context"
	"log"

	pb "github.com/Code2aum/grpc-demo/proto" // Import your generated protobuf package
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCurrencyServiceClient(conn)

	// Example usage:
	// Create currency
	_, err = c.CreateCurrency(context.Background(), &pb.Currency{Name: "USD", Value: 1.0})
	if err != nil {
		log.Fatalf("CreateCurrency failed: %v", err)
	}

	// Read currency
	readResp, err := c.ReadCurrency(context.Background(), &pb.ReadCurrencyRequest{Name: "USD"})
	if err != nil {
		log.Fatalf("ReadCurrency failed: %v", err)
	}
	log.Printf("Read currency: %s - %f", readResp.Name, readResp.Value)

	// Update currency
	_, err = c.UpdateCurrency(context.Background(), &pb.Currency{Name: "USD", Value: 1.2})
	if err != nil {
		log.Fatalf("UpdateCurrency failed: %v", err)
	}

	// Delete currency
	_, err = c.DeleteCurrency(context.Background(), &pb.DeleteCurrencyRequest{Name: "USD"})
	if err != nil {
		log.Fatalf("DeleteCurrency failed: %v", err)
	}

	// List currencies
	stream, err := c.ListCurrencies(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("ListCurrencies failed: %v", err)
	}
	for {
		currency, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("List currency: %s - %f", currency.Name, currency.Value)
	}
}
