package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float64) {
	fmt.Println("Payment made for", amount)

	bkashPayment := bkash{}
	bkashPayment.pay(amount)


}

type bkash struct {}

func (b bkash) pay(amount float64) {
	fmt.Println("Bkash payment made for", amount)
}

func main(){
	newPayment := payment{}
	newPayment.makePayment(100)

}