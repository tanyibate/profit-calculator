package main

import (
	"fmt"

	"example.com/profit-calculator/util"
)

func main() {
	var revenue, expenses float64
	const taxRate = 25
	fmt.Print("revenue:")
	fmt.Scan(&revenue)
	fmt.Print("expenses:")
	fmt.Scan(&expenses)

	earningsBeforeTax, earningsAfterTax, ratio := util.CalculateEarnings(revenue, expenses, taxRate)

	fmt.Print(earningsBeforeTax, earningsAfterTax, ratio)

}
