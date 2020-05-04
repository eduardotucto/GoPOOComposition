package invoice

import (
	"github.com/eduardotucto/GoPOOComposition/pkg/customer"
	"github.com/eduardotucto/GoPOOComposition/pkg/invoiceitem"
)

// Invoice es la estructura de una factura
// como se dice a Go que factura tiene relaion con cliente, luego tambien con item
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer  // 1 a 1
	items   []invoiceitem.Item // un invoice tiene n items
}

// SetTotal es el setter de Invoice.total
func (i *Invoice) SetTotal() {
	// recorre la cantidad de datos que tenga items
	// item1 alamacena el valor de un item
	// luego se extrae su precio con el getter value() de invoiceitem
	for _, item1 := range i.items {
		i.total += item1.Value()
	}
}
