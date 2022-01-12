package http

import (
	"context"
	"encoding/json"
	"net/http"

	"customerapp/services/customer"
	"customerapp/services/customer/transport"
)

func decodeCreateCustomerRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Customer); e != nil {
		return nil, e
	}
	return req, nil
}



func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case customer.ErrOrderNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
