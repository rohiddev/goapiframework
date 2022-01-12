package middleware


import "customerapp/services/customer"

// Middleware describes a service middleware.
type Middleware func(service customer.Service) customer.Service

