package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Number of expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Taces to account for: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses

	taxes := EBT * (1 - taxRate/100)

	PAT := EBT - taxes

	ratio := EBT / PAT

	fmt.Print("EBT before taxes: ")
	fmt.Println(EBT)

	fmt.Print("Profit after taxes: ")
	fmt.Println(PAT)

	fmt.Print("Ratio ebt/profit: ")
	fmt.Println(ratio)
}
