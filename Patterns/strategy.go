package Patterns

import (
	"fmt"
	"strconv"
)

func present() {
	var a string
	product := "phone"
	fmt.Scanln(&a)
	t, _ := strconv.Atoi(a)
	var payment Payment
	switch t {
	case 1:
		payment = NewCardPayment("1233", "124")
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()

	}

	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	err := payment.Pay()
	if err != nil {
		return
	}
}

type Payment interface {
	Pay() error
}

type cardPayment struct {
	cardNumber, cvv string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *cardPayment) Pay() error {
	return nil
}

type payPalPayment struct {
}

func NewPayPalPayment() Payment {
	return &payPalPayment{}
}

func (p *payPalPayment) Pay() error {
	// ... implementation
	return nil
}

type qiwiPayment struct {
}

func NewQIWIPayment() Payment {
	return &qiwiPayment{}
}

func (p *qiwiPayment) Pay() error {
	return nil
}
