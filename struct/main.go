package main

import (
	"fmt"
	"time"
)

// create struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// constructor
func newOrderC(id string, amount float32, status string) *order {
	return &order{
		id:     id,
		amount: amount,
		status: status,
	}
}

// reciever reference of instance as its receiver has pointer to blueprint of order struct
func (o *order) changeStatus(s string) {
	o.status = s
}

func main() {
	// very basic struct instance
	// default value is set to zero if not assigned depending on its type
	newOrder := order{
		id:     "1",
		amount: 99.0,
		status: "created",
	}
	newOrder.createdAt = time.Now().UTC()

	// using method to change status
	// its always a good practice to create methods for operations
	newOrder.changeStatus("delivered")

	newOrderAgain := order{
		id:        "2",
		amount:    99.0,
		status:    "shipped",
		createdAt: time.Now(),
	}

	fmt.Println(newOrder)
	fmt.Println(newOrderAgain)

	// using constuctor
	newOrderThroughC := newOrderC("1", 213.00, "created")
	fmt.Println(newOrderThroughC.amount)

	// structs can be crated for one time use only as well

	l:=struct{
		name string
	}{"go"}

	fmt.Println(l)
}
