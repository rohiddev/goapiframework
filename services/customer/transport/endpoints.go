package transport

import (
	"context"
	"customerapp/services/customer"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds all Go kit endpoints for the Order service.
type Endpoints struct {
	CreateCustomer       endpoint.Endpoint

}

// MakeEndpoints initializes all Go kit endpoints for the Order service.
func MakeEndpoints(s customer.Service) Endpoints {
	return Endpoints{
		CreateCustomer:       makeCreateEndpoint(s),

	}
}


func makeCreateEndpoint(s customer.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest) // type assertion
		id, err := s.CreateCustomer(ctx, req.Customer)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

