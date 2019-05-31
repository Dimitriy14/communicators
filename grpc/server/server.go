package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Dimitriy14/communicators/grpc/communicatorpb"
)

type server struct{}

func (*server) GetAvgSlice(ctx context.Context, req *communicatorpb.ProductSliceRequest) (*communicatorpb.ProductResponse, error) {
	sum := float32(0)
	amountSum := int32(0)

	for _, product := range req.Product {
		sum += product.GetPrice() * float32(product.GetAmount())

		amountSum += product.GetAmount()
	}

	return &communicatorpb.ProductResponse{Avg: sum / float32(amountSum)}, nil
}

//func multiply(a float32, b int32) float32 {
//	return a * float32(b)
//}

func (*server) GetAvg(stream communicatorpb.AvgService_GetAvgServer) error {
	//fmt.Println("GetAvg started...")

	sum := float32(0)
	amountSum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := sum / float32(amountSum)
			return stream.SendAndClose(&communicatorpb.ProductResponse{Avg: average})
		}
		if err != nil {
			log.Fatal(err)
		}
		sum += req.GetProduct().GetPrice() * float32(req.GetProduct().GetAmount())
		amountSum += req.GetProduct().GetAmount()
	}
}

func main() {
	//fmt.Println("Average Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	communicatorpb.RegisterAvgServiceServer(s, &server{})

	fmt.Println("Start Serving...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
