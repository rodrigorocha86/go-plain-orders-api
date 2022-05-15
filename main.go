package main

import (
	"fmt"
	"go-plain-orders-domain/customer"
)

func main() {
	customer := customer.NewCustomer("001", customer.NewCpf("55555555555"), "Fulano")
	fmt.Printf("Code: %v and Name: %v\n", customer.GetCode(), customer.GetName())

	customer.ChangeName("Beltrano")
	fmt.Printf("Code: %v and Name: %v\n", customer.GetCode(), customer.GetName())

	customers := CreateCustomers()
	customers.AddCustomer(customer)
	fmt.Println(customers.GetCustomer("001"))
	fmt.Println(customers.GetCustomer("002"))

	customer.ChangeName("Ciclano")
	customers.UpdateCustomer(customer)
	fmt.Println(customers.GetCustomer("001"))
}