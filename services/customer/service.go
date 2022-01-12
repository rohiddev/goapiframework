package customer



import (
	"context"
	"errors"
)

var (
	ErrOrderNotFound   = errors.New("order not found")
	ErrCmdRepository   = errors.New("unable to command repository")
	ErrQueryRepository = errors.New("unable to query repository")
)

// Service describes the Order service.
type Service interface {
	//new work
	CreateCustomer(ctx context.Context, customer Customer) (string, error)

}


