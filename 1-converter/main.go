package main

import "fmt"

const (
	DOLLARS_TO_EURO = 0.92
	DOLLARS_TO_RUB  = 92.50
)


func enterUserInput() {
	fmt.Println("Enter amount in USD:")
	var amount float64
	fmt.Scan(&amount)
}

func converCurrency(amount float64, origin string, target string) {}

func main() {
	euroToRub := DOLLARS_TO_RUB / DOLLARS_TO_EURO

	fmt.Printf("USD to EUR %.2f\n", DOLLARS_TO_EURO)
	fmt.Printf("USD to RUB %.2f\n", DOLLARS_TO_RUB)
	fmt.Printf("EUR to RUB %.2f\n", euroToRub)

	enterUserInput()
}

