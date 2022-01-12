package transport



import (
	"customerapp/services/customer"
)

// CreateRequest holds the request parameters for the Create method.
type CreateRequest struct {
	Customer customer.Customer
}

// CreateResponse holds the response values for the Create method.
type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

/*
To include business error messages as annotations
in OpenCensus spans we need the Go kit Response
structs to implement the endpoint.Failer interface.
An issue report has been filed so this next step
might become deprecated in the (near) future.
*/
// Failed implements Failer
func (r CreateResponse) Failed() error { return r.Err }
