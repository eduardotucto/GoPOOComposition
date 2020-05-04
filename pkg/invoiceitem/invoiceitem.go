package invoiceitem

// Item contiene informacion de un invoiceitem(factura item)
type Item struct {
	id      uint
	product string
	value   float64
}

// New retorna un nuevo Item (constructor)
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value obtiene valor de value
func (i Item) Value() float64 {
	return i.value
}
