package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

type customer struct {
	id   string
	name string
}

func newOrder(id string, amount float32, status string) *order {
	return &order{
		id:     id,
		amount: amount,
		status: status,
	}
}

func newCustomer(id string, name string) *customer {
	return &customer{
		id:   id,
		name: name,
	}
}

// composition and inheritance can be achieved
func main() {
	order1 := newOrder("1", 42.0, "created")
	custome1 := newCustomer("1", "gaurav")

	order1.customer = *custome1

	fmt.Println(order1)

}
