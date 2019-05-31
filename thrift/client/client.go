package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"log"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/Dimitriy14/communicators/thrift/gen-go/communicator"
)

func main() {
	time.Sleep(time.Second * 10)
	client, transport, err := NewThriftClient("server:1488")

	if err != nil {
		log.Fatal(err)
	}

	defer transport.Close()

	var sum time.Duration

	for i := 0; i < 10000; i++ {
		t1 := time.Now()

		getAvg(client)

		t2 := time.Now()

		sum += t2.Sub(t1)
	}

	fmt.Println(sum / 10000)
}

func NewThriftClient(url string) (*communicator.AvgServiceClient, *thrift.TSocket, error) {
	transport, err := thrift.NewTSocket(url)

	if err != nil {
		return nil, nil, fmt.Errorf("transport error: %v", err)
	}

	if err = transport.Open(); err != nil {
		log.Fatal(err)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	useTransport, _ := transportFactory.GetTransport(transport)

	client := communicator.NewAvgServiceClientFactory(useTransport, protocolFactory)

	return client, transport, nil
}

func getAvg(client *communicator.AvgServiceClient) {
	file, err := os.Open("average.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	var products []*communicator.Product

	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		price, _ := strconv.ParseFloat(line[1], 32)
		amount, _ := strconv.ParseInt(line[2], 10, 32)

		products = append(products, &communicator.Product{Price: price, Amount: int32(amount)})
	}

	_, err = client.GetAvg(context.Background(), products)

	if err != nil {
		log.Fatal(err)
	}
}
