package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggerService interface {
	FetchPrice(context.Context, string) (float64, error)
}

type loggerService struct {
	next PriceFetcher
}

func NewLoggerService(next PriceFetcher) PriceFetcher {
	return &loggerService{
		next: next,
	}
}

func (ls *loggerService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(startTime time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"time":      time.Since(startTime),
			"price":     price,
			"error":     err,
		}).Info("FETCH PRICE")
	}(time.Now())
	fmt.Println("FETCH")

	return ls.next.FetchPrice(ctx, ticker)
}
