package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr",":3000","the listen address the service is running")
	flag.Parse()

	svc := NewMetricService(NewLoggingService(&priceFetcher{}))

	server := NewJSONApiServer(*listenAddr,svc)
	server.Run()
}