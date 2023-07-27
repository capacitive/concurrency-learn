package main

import "fmt"

type order struct {
	ProductCode int
	Quantity    float64
	Status      OrderStatus
}

type invalidOrder struct {
	order order
	err   error
}

func (o order) String() string {
	return fmt.Sprintf("Product code: %v, Quantity: %v, Status: %v\n",
		o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

// func (o OrderStatus) String() string {
// 	return [...]string{"none", "new", "received", "reserved", "filled", "unknown"}[o]
// }

func orderStatusToText(o OrderStatus) string {
	switch o {
	case none:
		return "none"
	case new:
		return "new"
	case received:
		return "received"
	case reserved:
		return "reserved"
	case filled:
		return "filled"
	default:
		return "unknown status"
	}
}

type OrderStatus int

const (
	none OrderStatus = iota
	new
	received
	reserved
	filled
	unknown
)

//var orders = []order{} //slice
