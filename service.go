package main

import (
	"context"
	"fmt"
	"time"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct {

}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx,ticker)
}

var priceMocks = map[string]float64{
	"BTC": 10000,
	"ETH": 1000,
	"XRP": 1,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	time.Sleep(time.Millisecond * 120)
	price,ok := priceMocks[ticker]
	if !ok {
		return price,fmt.Errorf("Ticker is not supported")
	}

	return price,nil
} 