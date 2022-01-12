package http

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"errors"
	"customerapp/services/customer/transport"
)



var (
	ErrBadRouting = errors.New("bad routing")
)

// initializeRoutes maps the Go kit endpoints
// to endpoints of HTTP request router
func initializeRoutes(svcEndpoints transport.Endpoints, options []kithttp.ServerOption) *mux.Router {
	r := mux.NewRouter()
	// HTTP Post - /orders
	r.Methods("POST").Path("/customers").Handler(kithttp.NewServer(
		svcEndpoints.CreateCustomer,
		decodeCreateCustomerRequest,
		encodeResponse,
		options...,
	))


	return r
}

