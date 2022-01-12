package implementation

import (
	"context"
	_ "database/sql"
	_ "time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"

	customersvc "customerapp/services/customer"
)

// service implements the Order Service
type service struct {
	repository customersvc.Repository
	logger     log.Logger
}

// NewService creates and returns a new Order service instance
func NewService(rep customersvc.Repository, logger log.Logger) customersvc.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// Create makes an order
func (s *service) CreateCustomer(ctx context.Context, customer customersvc.Customer) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	customer.ID = id


	if err := s.repository.CreateCustomer(ctx, customer); err != nil {
		level.Error(logger).Log("err", err)
		return "", customersvc.ErrCmdRepository
	}
	return id, nil
}

