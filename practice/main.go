package main

import "fmt"

type PaymentProvider interface {
	pay(amount int) error
	name() string
}

type Razorpay struct {
	apiKey    string
	apiSecret string
}

func (r *Razorpay) pay(amount int) error {
	return nil
}

func (r *Razorpay) name() string {
	return "Razorpay"
}

type Stripe struct {
	apiKey    string
	apiSecret string
}

func (s *Stripe) pay(amount int) error {
	return nil
}

func (s *Stripe) name() string {
	return "Stripe"
}

type PaymentSystem interface {
	processPayment(amount int) error
}

type PaymentProcessor struct {
	provider PaymentProvider
}

func (p *PaymentProcessor) processPayment(amount int) error {
	fmt.Println("Processing payment of", amount)
	fmt.Println("Using provider:", p.provider.name())
	return p.provider.pay(amount)
}

func main() {
	razorpay := &Razorpay{
		apiKey:    "key",
		apiSecret: "secret",
	}

	processor := PaymentProcessor{
		provider: razorpay,
	}

	err := processor.processPayment(1000)
	if err != nil {
		fmt.Println(err)
	}
}
