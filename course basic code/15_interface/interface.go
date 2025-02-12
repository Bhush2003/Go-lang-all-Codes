package main

import "fmt"

// interface
type paymenter interface {
	pay(amount int)
}

// struct for payment
type makepay struct{
	getway paymenter
}

type paypal struct{}

func (p paypal) pay(amount int) {
	fmt.Println("paypal",amount)
}
type rozerpay struct{}

func (p rozerpay) pay(amount int) {
	fmt.Println("rozerpay",amount)
}

func (m makepay) makePayment(amount int){
	m.getway.pay(amount)
}

func main() {
	// paymentGW:=rozerpay{}
	// paymentGW1:=paypal{}
	make:=makepay{
		getway: paypal{},
	}
	make1:=makepay{
		getway: rozerpay{},
	}
	make.makePayment(10)
	make1.makePayment(40)

	
}

// interfaces in go are only dependent on method signeture and not on any anather method