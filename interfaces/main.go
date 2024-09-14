package main

import (
	"fmt"
)

type PaymentI interface {
	pay(a float32)
}

type razorpay struct{}
type stripe struct{}

type payment struct {
	gateway PaymentI
}

func (r razorpay) pay(a float32) {
	fmt.Println("razorpay is being used", a)
}

func (s stripe) pay(a float32) {
	fmt.Println("stripe is being used", a)
}

func (p payment) makePayment(a float32) {
	p.gateway.pay(a)
}

func main() {
	// rpg := razorpay{}
	// newPayment := payment{
	// 	gateway: rpg,
	// }

	// newPayment.makePayment(30)
	// s := notifier.Sms{}
	// n := Notifier{
	// 	notifier: s,
	// }
	// n.Notification([]string{"gaurav", "alice"}, "you did it")
}
