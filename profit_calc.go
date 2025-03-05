package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&revenue)

	//fmt.Print("Number of expenses: ")
	outputText("Number of expenses: ")
	fmt.Scan(&expenses)

	//fmt.Print("Taxes to account for: ")
	outputText("Taxes to account for: ")
	fmt.Scan(&taxRate)

	//EBT := revenue - expenses

	//taxes := EBT * (1 - taxRate/100)

	//PAT := EBT - taxes

	//ratio := EBT / PAT

	EBTCalc, PATCalc, ratioCalc := calculating(revenue, expenses, taxRate)

	//fmt.Print("EBT before taxes: ")
	outputText("EBT before taxes: ")
	fmt.Println(EBTCalc)

	//fmt.Print("Profit after taxes: ")
	outputText("Profit after taxes: ")
	fmt.Println(PATCalc)

	//fmt.Print("Ratio ebt/profit: ")
	outputText("Ratio ebt/profit: ")
	fmt.Println(ratioCalc)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculating(revenue, expenses, taxRate float64) (float64, float64, float64) {
	EBT := revenue - expenses
	taxes := EBT * (1 - taxRate/100)
	PAT := EBT - taxes
	ratio := EBT / PAT
	return EBT, PAT, ratio
}
