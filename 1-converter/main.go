package main

import (
	"fmt"
)

const (
	DOLLARS_TO_EURO = 0.92
	DOLLARS_TO_RUB  = 92.50
)

var currencies = []string{"USD", "EUR", "RUB"}

func promptCurrency(prompt string, options []string) string {
	for {
		var opts string
		for i, c := range options {
			if i > 0 {
				opts += "/"
			}
			opts += c
		}
		fmt.Printf("%s (%s): ", prompt, opts)
		var input string
		fmt.Scan(&input)
		if input == "usd" || input == "Usd" {
			input = "USD"
		}
		if input == "eur" || input == "Eur" {
			input = "EUR"
		}
		if input == "rub" || input == "Rub" {
			input = "RUB"
		}
		for _, c := range options {
			if input == c {
				return input
			}
		}
		fmt.Println("Invalid currency. Try again.")
	}
}

func promptAmount() float64 {
	for {
		fmt.Print("Enter amount: ")
		var amount float64
		_, err := fmt.Scan(&amount)
		if err == nil && amount >= 0 {
			return amount
		}
		fmt.Println("Invalid amount. Try again.")
	}
}

func convertCurrency(amount float64, origin, target string) float64 {
	switch origin + "_" + target {
	case "USD_EUR":
		return amount * DOLLARS_TO_EURO
	case "USD_RUB":
		return amount * DOLLARS_TO_RUB
	case "EUR_USD":
		return amount / DOLLARS_TO_EURO
	case "EUR_RUB":
		return amount * (DOLLARS_TO_RUB / DOLLARS_TO_EURO)
	case "RUB_USD":
		return amount / DOLLARS_TO_RUB
	case "RUB_EUR":
		return amount / (DOLLARS_TO_RUB / DOLLARS_TO_EURO)
	default:
		return amount
	}
}

func filterCurrencies(exclude string) []string {
	var filtered []string
	for _, c := range currencies {
		if c != exclude {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func main() {
	origin := promptCurrency("Enter origin currency", currencies)
	amount := promptAmount()
	targetOptions := filterCurrencies(origin)
	target := promptCurrency("Enter target currency", targetOptions)

	result := convertCurrency(amount, origin, target)
	fmt.Printf("%.2f %s = %.2f %s\n", amount, origin, result, target)
}
