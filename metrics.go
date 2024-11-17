package main

import (
	"context"
	"fmt"
)

type MetricsService interface {
	FetchPrice(context.Context, string) (float64, error)
}

type metricsService struct {
	next PriceFetcher
}

func NewMetricsService(next PriceFetcher) PriceFetcher {
	return &metricsService{
		next: next,
	}
}

func (svc *metricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func() {
		fmt.Printf("Metrics price for %s", ticker)
	}()

	return svc.next.FetchPrice(ctx, ticker)
}
