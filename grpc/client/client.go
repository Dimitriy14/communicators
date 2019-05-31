package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/Dimitriy14/communicators/grpc/communicatorpb"
)

func main() {
	time.Sleep(time.Second * 10)
	cc, err := grpc.Dial("server:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	c := communicatorpb.NewAvgServiceClient(cc)

	//CalculateAvg(c)
	var sum time.Duration

	for i := 0; i < 10000; i++ {
		t1 := time.Now()
		CalculateAvgFromSlice(c)
		t2 := time.Now()

		sum += t2.Sub(t1)
	}

	fmt.Println(sum / 10000)
}

func CalculateAvg(c communicatorpb.AvgServiceClient) {
	stream, err := c.GetAvg(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("average.csv")

	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

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

		if err = stream.Send(&communicatorpb.ProductRequest{Product: &communicatorpb.Product{Price: float32(price), Amount: int32(amount)}}); err != nil {
			log.Fatal(err)
		}
	}

	_, err = stream.CloseAndRecv()

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Result: %v\n", res.GetAvg())
}

func CalculateAvgFromSlice(c communicatorpb.AvgServiceClient) {
	file, err := os.Open("average.csv")

	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	var products []*communicatorpb.Product
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

		products = append(products, &communicatorpb.Product{Price: float32(price), Amount: int32(amount)})
	}

	_, err = c.GetAvgSlice(context.Background(), &communicatorpb.ProductSliceRequest{Product: products})

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("RESULT: %f\n", resp.Avg)
}
