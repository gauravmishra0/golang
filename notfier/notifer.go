package main

import (
	"fmt"
)

// TODO: implement ratelimiting, goroutine to send notification on different channel, bulk method,
// retry mech, circuit Breaker, Queue-based architecture and Service Registry

// Circuit Breaker Pattern:
// Problem: If a notification service (e.g., email, SMS) fails repeatedly, continuing to try could degrade the system.
// Optimization: Implement a circuit breaker (using a library like go-resilience) to stop sending requests to a failed
// service temporarily and automatically retry after a cool-down period.
// Example:
// breaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{})
// _, err := breaker.Execute(func() (interface{}, error) {
//     return actualSendMail(to, message)
// })

// NotifierI defines the interface for all notifiers
type NotifierI interface {
	Send(to []string, message string) error
}

// MultiNotifier allows sending notifications to multiple channels
type MultiNotifier struct {
	notifiers []NotifierI
}

// Sms, InApp, and Mail are concrete implementations of NotifierI
type Sms struct{}
type InApp struct{}
type Mail struct{}

// NewNotifier creates a multi-channel notifier
func NewNotifier(n ...NotifierI) MultiNotifier {
	return MultiNotifier{
		notifiers: n,
	}
}

// Notification sends messages to all channels
func (n MultiNotifier) Notification(to []string, message string) {
	for _, notifier := range n.notifiers {
		err := notifier.Send(to, message)
		if err != nil {
			fmt.Printf("failed to send notification via %T: %v\n", notifier, err)
		}
	}
}

// Send methods for different notifiers
func (s Sms) Send(to []string, message string) error {
	fmt.Printf("sending sms to %v : %s\n", to, message)
	return nil
}

func (i InApp) Send(to []string, message string) error {
	fmt.Printf("sending inapp to %v : %s\n", to, message)
	return nil
}

func (m Mail) Send(to []string, message string) error {
	fmt.Printf("sending mail to %v : %s\n", to, message)
	return nil
}

func main() {
	sms := Sms{}
	mail := Mail{}
	inApp := InApp{}

	// Create a multi-channel notifier
	multiNotifier := NewNotifier(sms, mail, inApp)

	// Send notifications to all channels
	multiNotifier.Notification([]string{"gaurav", "alice"}, "you did it")
}
