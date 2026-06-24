// Design a payment processing system that supports multiple payment gateways and notification channels 
// using interfaces and dependency injection.
// The system should validate user input, process payments, and send success or failure notifications accordingly.

package main

import (
	"errors"
	"fmt"
)

type PaymentMetadata struct {
	amount  string
	name    string
	email   string
	phone   string
	city    string
	country string
}

type PaymentGateway interface {
	ProcessPayment(metadata PaymentMetadata) error
}

type Razorpay struct{}

func (r *Razorpay) ProcessPayment(metadata PaymentMetadata) error {
	if metadata.amount == "0" {
		return errors.New("amount is required")
	}
	fmt.Println("Processing payment with Razorpay")
	return nil
}

type Stripe struct{}

func (s *Stripe) ProcessPayment(metadata PaymentMetadata) error {
	fmt.Println("Processing payment with Stripe")
	return nil
}

type Notifications interface {
	SendNotification(message string) error
}

type EmailNotifications struct{}

func (e *EmailNotifications) SendNotification(message string) error {
	fmt.Println("Sending email notification:", message)
	return nil
}

type SMSNotifications struct{}

func (s *SMSNotifications) SendNotification(message string) error {
	fmt.Println("Sending SMS notification:", message)
	return nil
}

type PaymentService struct {
	gateway       PaymentGateway
	notifications []Notifications
}

func (p *PaymentService) Pay(metadata PaymentMetadata) error {
	if metadata.amount == "" || metadata.amount == "0" {
		return errors.New("amount is required")
	}
	if metadata.name == "" {
		return errors.New("name is required")
	}
	if metadata.email == "" {
		return errors.New("email is required")
	}
	if metadata.phone == "" {
		return errors.New("phone is required")
	}
	err := p.gateway.ProcessPayment(metadata)
	if err != nil {
		for _, notification := range p.notifications {
			notification.SendNotification("Payment failed")
		}
		return err
	}
	for _, notification := range p.notifications {
		notification.SendNotification("Payment successful")
	}
	return nil
}

func main() {
	razorpay := &Razorpay{}
	// stripe := &Stripe{} // You can add more payment gateways here
	paymentService := &PaymentService{
		gateway:       razorpay,
		notifications: []Notifications{&EmailNotifications{}, &SMSNotifications{}},
	}
	err := paymentService.Pay(PaymentMetadata{
		amount:  "1",
		name:    "Ayush",
		email:   "ayush4002gupta@gmail.com",
		phone:   "992929299291",
		city:    "New York",
		country: "USA",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}
