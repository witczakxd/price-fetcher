package main

import (
	//"context"
	"flag"
	//"fmt"
	//"log"

	//"github.com/witczakxd/price-fetcher/client"
)

func main() {
	//client := client.New("http://localhost:3000")
	//price ,err := client.FetchPrice(context.Background(),"ET")
	//if err != nil {
	//	log.Fatal(err)
	//}
//
	//fmt.Printf("%+v\n",price)
	//return
	listenAddr := flag.String("listenaddr",":3000","the listen address the service is running")
	flag.Parse()

	svc := NewMetricService(NewLoggingService(&priceFetcher{}))

	server := NewJSONApiServer(*listenAddr,svc)
	server.Run()
}