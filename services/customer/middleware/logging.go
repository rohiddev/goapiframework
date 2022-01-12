package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"customerapp/services/customer"
)

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next customer.Service) customer.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   customer.Service
	logger log.Logger
}



func (mw loggingMiddleware) CreateCustomer(ctx context.Context, customer customer.Customer) (id string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Create", "CustomerID", customer.CustomerID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.CreateCustomer(ctx, customer)
}

