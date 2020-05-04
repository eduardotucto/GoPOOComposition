package main

import (
	"fmt"

	"github.com/eduardotucto/GoPOOComposition/pkg/customer"
	"github.com/eduardotucto/GoPOOComposition/pkg/invoice"
	"github.com/eduardotucto/GoPOOComposition/pkg/invoiceitem"
)

func main() {
	// crearemos una factura
	i := invoice.New(
		"Colombia",
		"Popayán",
		customer.New("Alejandro", "Cl 123 #234", "12345"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Curso de Go!", 12.34),
			invoiceitem.New(2, "Curso de POO con Go!", 54.23),
			invoiceitem.New(3, "Curso de testing con Go!", 90.00),
		},
	)
	i.SetTotal()         // halla el total 156.57
	fmt.Printf("%+v", i) // imprime todos los campos de invoice
}
