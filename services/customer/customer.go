package customer

import "context"



type Customer struct {
	ID           string      `json:"id,omitempty"`
	CustomerID   string      `json:"customer_id"`
	Address      string      `json:"address,omitempty"`

}

// Repository describes the persistence on order model
type Repository interface {


	// - added customer details
	CreateCustomer(ctx context.Context, customer Customer) error
	//GetCustomerByID(ctx context.Context, id string) (Customer, error)

}

