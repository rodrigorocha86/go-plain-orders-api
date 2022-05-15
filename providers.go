//go:build wireinject
package main

import (
	"github.com/google/wire"
	"go-plain-orders-domain/customer"
)

func CreateCustomers() customer.Customers {
	wire.Build(newCustomers)
	return *new(customer.Customers)
}

func newCustomers() customer.Customers {
	return inMemoryCustomers{}
}