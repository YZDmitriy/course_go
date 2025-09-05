package main

import (
    "fmt"
)

var rates = map[string]float64{
    "USD": 1.0,
    "EUR": 0.92,
    "RUB": 92.50,
}

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
        input = normalizeCurrency(input)
        for _, c := range options {
            if input == c {
                return input
            }
        }
        fmt.Println("Invalid currency. Try again.")
    }
}

func normalizeCurrency(input string) string {
    switch input {
    case "usd", "Usd":
        return "USD"
    case "eur", "Eur":
        return "EUR"
    case "rub", "Rub":
        return "RUB"
    default:
        return input
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

func convertCurrency(amount float64, origin, target string, rates *map[string]float64) float64 {
    originRate, ok1 := (*rates)[origin]
    targetRate, ok2 := (*rates)[target]
    if !ok1 || !ok2 {
        return amount
    }
    usdAmount := amount / originRate
    return usdAmount * targetRate
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

    result := convertCurrency(amount, origin, target, &rates)
    fmt.Printf("%.2f %s = %.2f %s\n", amount, origin, result, target)
}
