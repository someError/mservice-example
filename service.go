package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (p *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return p.MockFetchPrice(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC":  80_000_000,
	"DASH": 2_000,
	"ATOM": 120,
}

func (p *priceFetcher) MockFetchPrice(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("Ticker %s is not supported", ticker)
	}

	return price, nil
}
