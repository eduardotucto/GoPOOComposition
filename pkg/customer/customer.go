package customer

// Customer es la estructura de cliente
type Customer struct {
	name    string
	address string
	phone   string
}

// New retorna un nuevo Customer
func New(name, address, phone string) Customer {
	return Customer{name, address, phone}
}
