package customer

type Customers interface {
	AddCustomer(customer Customer)
	UpdateCustomer(customer Customer)
	GetCustomer(code string) Customer
}