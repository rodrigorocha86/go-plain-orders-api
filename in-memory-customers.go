package main

import "go-plain-orders-domain/customer"

var data = make(map[string]customer.Customer)

type inMemoryCustomers struct {}

func (inMemoryCustomers) AddCustomer(customer customer.Customer) {
	data[customer.GetCode()] = customer
}

func (inMemoryCustomers) UpdateCustomer(customer customer.Customer) {
	data[customer.GetCode()] = customer
}

func (inMemoryCustomers) GetCustomer(code string) customer.Customer {
	return data[code]
}