package customer

type Customer struct {
    code string
    cpf  Cpf
    name string
}

func New(code string, cpf Cpf, name string) *Customer {
    return &Customer{code: code, cpf: cpf, name: name}
}

func (c *Customer) GetCode() string {
    return c.code
}

func (c *Customer) GetName() string {
    return c.name
}

func (c *Customer) ChangeName(newName string) {
    c.name = newName
}
