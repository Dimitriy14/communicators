package main

import (
	"context"

	"log"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/Dimitriy14/communicators/thrift/gen-go/communicator"
)

type server struct {
}

func (*server) GetAvg(ctx context.Context, products []*communicator.Product) (r float64, err error) {
	sum := float64(0)
	amountSum := int32(0)

	for _, product := range products {
		sum += product.GetPrice() * float64(product.GetAmount())

		amountSum += product.GetAmount()
	}

	return sum / float64(amountSum), nil
}

func main() {
	serverTransport, err := thrift.NewTServerSocket("server:1488")

	if err != nil {
		log.Fatal(err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	handler := &server{}

	processor := communicator.NewAvgServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, thrift.NewTBinaryProtocolFactoryDefault())

	log.Fatal(server.Serve())

	defer server.Stop()
}
