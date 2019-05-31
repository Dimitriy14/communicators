package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main(){
	file, err:= os.Create("average.csv")

	if err!=nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer:= csv.NewWriter(file)
	defer writer.Flush()

	for i:= 0; i < 10000; i++{
		if err = writer.Write([]string{fmt.Sprintf("%v", i), fmt.Sprintf("%.2f", rand.ExpFloat64()* 100), fmt.Sprintf("%v", rand.Int31()%1000)}); err!=nil{
			log.Fatal(err)
		}
	}
}
