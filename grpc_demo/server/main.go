package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Code2aum/grpc-demo/proto" // Import your generated protobuf package
	"google.golang.org/grpc"
)

// CurrencyServer implements the CurrencyServiceServer interface
type CurrencyServer struct {
	pb.UnimplementedCurrencyServiceServer // Embedding the unimplemented struct to satisfy the interface
	currencies                            map[string]float32
}

// NewCurrencyServer creates a new instance of CurrencyServer
func NewCurrencyServer() *CurrencyServer {
	return &CurrencyServer{
		currencies: make(map[string]float32),
	}
}

// CreateCurrency creates a new currency
func (s *CurrencyServer) CreateCurrency(ctx context.Context, in *pb.Currency) (*pb.Currency, error) {
	s.currencies[in.Name] = in.Value
	return in, nil
}

// ReadCurrency reads a currency
func (s *CurrencyServer) ReadCurrency(ctx context.Context, in *pb.ReadCurrencyRequest) (*pb.Currency, error) {
	value, ok := s.currencies[in.Name]
	if !ok {
		return nil, fmt.Errorf("Currency %s not found", in.Name)
	}
	return &pb.Currency{Name: in.Name, Value: value}, nil
}

// UpdateCurrency updates a currency
func (s *CurrencyServer) UpdateCurrency(ctx context.Context, in *pb.Currency) (*pb.Currency, error) {
	if _, ok := s.currencies[in.Name]; !ok {
		return nil, fmt.Errorf("Currency %s not found", in.Name)
	}
	s.currencies[in.Name] = in.Value
	return in, nil
}

// DeleteCurrency deletes a currency
func (s *CurrencyServer) DeleteCurrency(ctx context.Context, in *pb.DeleteCurrencyRequest) (*pb.DeleteCurrencyResponse, error) {
	if _, ok := s.currencies[in.Name]; !ok {
		return nil, fmt.Errorf("Currency %s not found", in.Name)
	}
	delete(s.currencies, in.Name)
	return &pb.DeleteCurrencyResponse{Success: true}, nil
}

// ListCurrencies lists all currencies
func (s *CurrencyServer) ListCurrencies(empty *pb.Empty, stream pb.CurrencyService_ListCurrenciesServer) error {
	for name, value := range s.currencies {
		if err := stream.Send(&pb.Currency{Name: name, Value: value}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCurrencyServiceServer(s, NewCurrencyServer())
	log.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
