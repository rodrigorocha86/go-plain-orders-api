package customer

type Cpf struct {
	value string
}

func NewCpf(value string) Cpf {
	return Cpf{value: value}
}
