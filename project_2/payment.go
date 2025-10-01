package main

import "fmt"

// we declared an interface for the method pay
type Payment interface{
	pay(amount int)
}

// declared three struct to different payment method
type Paypal struct{}
type Stripe struct{}
type Crypto struct{}

// function to show the different payment method
func (p Paypal) pay(amount int){
	fmt.Printf("Amount paid %d is done through paypal\n",amount)
}
func (p Stripe) pay(amount int){
	fmt.Printf("Amount paid %d is done through Stripe\n",amount)
}
func (p Crypto) pay(amount int){
	fmt.Printf("Amount paid %d is done through Crypto\n",amount)
}

func main(){
	// storing all the struct together
	Payments := []Payment{Paypal{},Stripe{},Crypto{}}

	for _ , Payment :=  range Payments {
		// Calling payment method one by one
		Payment.pay(100)
	}
}